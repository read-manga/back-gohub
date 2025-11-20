package usecase

import (
	"microgo/core/domain/change"
	"microgo/infrastructure/repository"
)

type ChangeCase struct {
	change repository.ChangeRepository
}

func NewChangeCase(c repository.ChangeRepository) *ChangeCase {
	return &ChangeCase{change: c}
}

func (c *ChangeCase) SaveChanges(ch change.Change) error {
	return c.change.Create(ch)
}

func (c *ChangeCase) GetChanges(commitId string) (*[]change.Change, error) {
	return c.change.FindAllChangeCommitId(commitId)
}
