package adapters

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"microgo/core/domain/change"

	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Adapter interface {
	UploadFileS3(bucket string, key string, filePath string) error

	CreateFileHash(b []byte) string

	DiffFolders(oldFolder *[]change.Change, newFolder []change.NewFilesBody) []change.Change
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

func (s *s3Adapter) CreateFileHash(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func (s *s3Adapter) DiffFolders(previous *[]change.Change, current []change.NewFilesBody) []change.Change {
	diffs := []change.Change{}

	prevMap := map[string]string{}
	currMap := map[string]string{}

	for _, p := range *previous {
		prevMap[p.FilePath] = p.NewHash
	}

	for _, c := range *current {
		currMap[c.Path] = c.Hash
	}

	for path, oldHash := range prevMap {
		newHash, exists := currMap[path]

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

	for path, oldHash := range currMap {
		if _, exitis := prevMap[path]; !exitis {
			diffs = append(diffs, change.Change{
				ChangeType:   change.Removed,
				PreviousHash: oldHash,
			})
		}
	}

	return diffs
}
