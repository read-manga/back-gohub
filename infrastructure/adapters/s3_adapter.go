package adapters

import (
	"context"
	"crypto/sha256"
	"fmt"
	"path/filepath"

	"microgo/core/domain/change"

	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Adapter interface {
	UploadFileS3(bucket string, key string, filePath string) error

	CreateFileHash(path string) (string, error)

	ScanFolder(basePath string) (map[string]string, error)

	DiffFolders(oldFolder string, newFolder string) ([]change.Change, error)
}

type s3Adapter struct{}

func NewS3Adapter() S3Adapter {
	return &s3Adapter{}
}

func (s *s3Adapter) UploadFileS3(bucket string, key string, filePath string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

func (s *s3Adapter) CreateFileHash(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(data)
	return fmt.Sprintf("%x", sum), nil
}

func (s *s3Adapter) ScanFolder(basePath string) (map[string]string, error) {
	result := make(map[string]string)

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(basePath, path)
		if err != nil {
			return err
		}

		hash, err := s.CreateFileHash(path)
		if err != nil {
			return err
		}

		result[rel] = hash
		return nil
	})

	return result, err
}

func (s *s3Adapter) DiffFolders(oldFolder string, newFolder string) ([]change.Change, error) {
	oldFiles, err := s.ScanFolder(oldFolder)
	if err != nil {
		return nil, err
	}

	newFiles, err := s.ScanFolder(newFolder)
	if err != nil {
		return nil, err
	}

	diffs := []change.Change{}

	for path, newHash := range newFiles {
		oldHash, exists := oldFiles[path]

		if !exists {
			diffs = append(diffs, change.Change{
				ChangeType: change.Added,
				NewHash:    newHash,
			})
			continue
		}

		if oldHash != newHash {
			diffs = append(diffs, change.Change{
				ChangeType:   change.Modified,
				PreviousHash: oldHash,
				NewHash:      newHash,
			})
		}
	}

	for path, oldHash := range oldFiles {
		if _, exitis := newFiles[path]; !exitis {
			diffs = append(diffs, change.Change{
				ChangeType:   change.Removed,
				PreviousHash: oldHash,
			})
		}
	}

	return diffs, nil

}
