package repository

import (
	"database/sql"
	"microgo/core/domain/repo"
)

type RepoRepository interface {
	Create(u repo.Repo) (string, error)
	FindAll(user_id string) (*[]repo.Repo, error)
	FindById(id string) (*repo.Repo, error)
	FindName(name string) (*[]repo.Repo, error)
}

type repoRepository struct {
	db *sql.DB
}

func NewRepoRepository(db *sql.DB) RepoRepository {
	return &repoRepository{db: db}
}

func (r *repoRepository) Create(entity repo.Repo) (string, error) {
	var id string

	err := r.db.QueryRow(`
		INSERT INTO repositorys (name, about, tag, public, storage_s3, user_id, colaborators)
        VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
		`,
		&entity.Name, &entity.About, &entity.Tag, &entity.Public, &entity.StorageS3, &entity.UserId, &entity.Colaborators,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil

}

func (r *repoRepository) FindAll(user_id string) (*[]repo.Repo, error) {
	query, err := r.db.Query(`SELECT * FROM repositorys WHERE user_id = $1`, user_id)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var repoRows []repo.Repo

	for query.Next() {
		var r repo.Repo

		err := query.Scan(&r.ID, &r.Name, &r.About, &r.Tag, &r.Stars, &r.Public)
		if err != nil {
			return nil, err
		}
		repoRows = append(repoRows, r)
	}

	err = query.Err()
	if err != nil {
		return nil, err
	}

	return &repoRows, nil
}

func (r *repoRepository) FindById(id string) (*repo.Repo, error) {
	var re = repo.Repo{}

	query := r.db.QueryRow(`SELECT * FROM repositorys WHERE id = $1`, id)

	err := query.Scan(&re.ID, &re.Name, &re.About, &re.Tag, &re.Stars, &re.Public, &re.StorageS3, &re.UserId, &re.Colaborators)
	if err != nil {
		return nil, err
	}

	return &re, nil
}

func (r *repoRepository) FindName(name string) (*[]repo.Repo, error) {
	query, err := r.db.Query(`SELECT * FROM repositorys WHERE name = $1`, name)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var repoRows []repo.Repo

	for query.Next() {
		var r repo.Repo

		err := query.Scan(&r.ID, &r.Name, &r.About, &r.Tag, &r.Stars, &r.Public, &r.UserId)
		if err != nil {
			return nil, err
		}

		repoRows = append(repoRows, r)
	}

	err = query.Err()
	if err != nil {
		return nil, err
	}

	return &repoRows, nil
}
