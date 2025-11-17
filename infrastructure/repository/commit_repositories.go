package repository

import (
	"database/sql"
	"microgo/core/domain/commit"
)

type CommitRepository interface {
	Create(u commit.Commit) error
	FindAllCommitRepoId(repo_id string) (*[]commit.Commit, error)
	FindCommitById(id string) (*commit.Commit, error)
}

type commitRepository struct {
	db *sql.DB
}

func NewCommitRepository(db *sql.DB) CommitRepository {
	return &commitRepository{db: db}
}

func (c *commitRepository) Create(entity commit.Commit) error {
	_, err := c.db.Exec(
		`INSERT INTO users (title, descritpion, repo_id, user_id, created_at)
         VALUES ($1, $2, $3, $4, $5)`,
		entity.Title, entity.Description, entity.RepoId, entity.UserId, entity.CreatedAt,
	)
	return err
}

func (c *commitRepository) FindAllCommitRepoId(repo_id string) (*[]commit.Commit, error) {
	query, err := c.db.Query(`SELECT * FROM commit WHERE repo_id = $1`, repo_id)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var commitRows []commit.Commit

	for query.Next() {
		var c commit.Commit

		err := query.Scan(&c.Title, &c.Description, &c.RepoId, &c.UserId, &c.CreatedAt)
		if err != nil {
			return nil, err
		}

		commitRows = append(commitRows, c)
	}

	err = query.Err()
	if err != nil {
		return nil, err
	}

	return &commitRows, nil
}

func (c *commitRepository) FindCommitById(id string) (*commit.Commit, error) {
	var co = commit.Commit{}

	query := c.db.QueryRow(`SELECT * FROM commit WHERE id = $1`, id)

	err := query.Scan(&co.ID, &co.Title, &co.Description, &co.RepoId, &co.UserId, &co.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &co, nil
}
