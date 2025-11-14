package usecase

import (
	"microgo/core/domain/repo"
	"microgo/infrastructure/repository"
)

type RepoUseCase struct {
	repo repository.RepoRepository
}

func NewRepoUseCase(r repository.RepoRepository) *RepoUseCase {
	return &RepoUseCase{repo: r}
}

func (r *RepoUseCase) Save(repo repo.Repo) error {
	return r.repo.Create(repo)
}
