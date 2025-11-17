package usecase

import (
	"errors"
	"microgo/core/domain/user"
	"microgo/infrastructure/repository"
	"microgo/infrastructure/security"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	user repository.UserRepository
}

func NewAuthUsecase(r repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{user: r}
}

func (a *AuthUsecase) Register(u user.User) error {
	exists, _ := a.user.FindByEmail(u.Email)
	if exists != nil {
		return errors.New("email já registrado")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return a.user.Create(u)
}

func (a *AuthUsecase) Login(email, password string) (string, error) {
	u, err := a.user.FindByEmail(email)
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

func (a *AuthUsecase) Update(data user.User) error {
	if data.ID == "" {
		return errors.New("id is required")
	}
	return a.user.Update(data)
}

func (a *AuthUsecase) Me(id string) (*user.User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	u, err := a.user.FindById(id)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.New("usuario nao encontrado")
	}

	return u, nil
}
