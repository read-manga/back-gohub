package repository

import (
	"database/sql"
	"microgo/core/domain/change"
)

type ChangeRepository interface {
	Create(entity change.Change) error
	FindAllChangeCommitId(commit_id string) (*[]change.Change, error)
}

type changeRepository struct {
	db *sql.DB
}

func NewChangeRepository(db *sql.DB) ChangeRepository {
	return &changeRepository{db: db}
}

func (c *changeRepository) Create(entity change.Change) error {
	_, err := c.db.Exec(
		`INSERT INTO change (commit_id, file_path, change_type, previous_hash, new_hash)
         VALUES ($1, $2, $3, $4, $5)`,
		entity.CommitId, entity.FilePath, entity.ChangeType, entity.PreviousHash, entity.NewHash,
	)
	return err
}

func (c *changeRepository) FindAllChangeCommitId(commit_id string) (*[]change.Change, error) {
	query, err := c.db.Query(`SELECT * FROM change WHERE commit_id=$1`, commit_id)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var changeRows []change.Change

	for query.Next() {
		var c change.Change

		err := query.Scan(&c.CommitId, &c.FilePath, &c.ChangeType, &c.PreviousHash, &c.NewHash)
		if err != nil {
			return nil, err
		}

		changeRows = append(changeRows, c)
	}

	err = query.Err()
	if err != nil {
		return nil, err
	}

	return &changeRows, nil
}
