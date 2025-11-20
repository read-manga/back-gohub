package repository

import (
	"context"
	"database/sql"
	"microgo/core/domain/commit"
)

type CommitRepository interface {
	Create(u commit.Commit) (string, error)
	FindAllCommitRepoId(repo_id string) (*[]commit.Commit, error)
	FindCommitById(id string) (*commit.Commit, error)
	FindCommitByDate(ctx context.Context) (*commit.Commit, error)
}

type commitRepository struct {
	db *sql.DB
}

func NewCommitRepository(db *sql.DB) CommitRepository {
	return &commitRepository{db: db}
}

func (c *commitRepository) Create(entity commit.Commit) (string, error) {
	var commitId string

	err := c.db.QueryRow(`INSERT INTO commit (title, descritpion, repo_id, user_id, created_at)
        VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		entity.Title, entity.Description, entity.RepoId, entity.UserId, entity.CreatedAt,
	).Scan(&commitId)

	if err != nil {
		return "", nil
	}

	return commitId, nil
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

func (c *commitRepository) FindCommitByDate(ctx context.Context) (*commit.Commit, error) {
	var commitQuery = commit.Commit{}

	query := c.db.QueryRowContext(ctx, `SELECT * FROM commit ORDER BY created_at DESC LIMIT 1;`)

	if err := query.Scan(&commitQuery.ID, &commitQuery.CreatedAt); err != nil {
		return nil, err
	}

	return &commitQuery, nil
}
