package usecase

import (
	"context"
	"microgo/core/domain/commit"
	"microgo/infrastructure/repository"
)

type CommitCase struct {
	commit repository.CommitRepository
}

func NewCommitUseCase(c repository.CommitRepository) *CommitCase {
	return &CommitCase{commit: c}
}

func (c *CommitCase) SaveCommit(co commit.Commit) (string, error) {
	return c.commit.Create(co)
}

func (c *CommitCase) GetCommitByDate() (*commit.Commit, error) {
	ctx := context.Background()
	return c.commit.FindCommitByDate(ctx)
}
