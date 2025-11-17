package usecase

import (
	"microgo/core/domain/commit"
	"microgo/infrastructure/repository"
)

type CommitCase struct {
	commit repository.CommitRepository
}

func NewCommitUseCase(c repository.CommitRepository) *CommitCase {
	return &CommitCase{commit: c}
}

func (c *CommitCase) SaveCommit(co commit.Commit) error {
	return c.commit.Create(co)
}
