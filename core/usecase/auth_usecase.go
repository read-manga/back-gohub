package usecase

import (
	"errors"
	"microgo/core/domain/user"
	"microgo/core/infra/security"
	"microgo/core/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(r repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{repo: r}
}

func (a *AuthUsecase) Register(u user.User) error {
	exists, _ := a.repo.FindByEmail(u.Email)
	if exists != nil {
		return errors.New("email já registrado")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return a.repo.Create(u)
}

func (a *AuthUsecase) Login(email, password string) (string, error) {
	u, err := a.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("usuário não encontrado")
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return "", errors.New("senha incorreta")
	}

	token, err := security.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
