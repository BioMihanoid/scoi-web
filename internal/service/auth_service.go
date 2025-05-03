package service

import (
	"golang.org/x/crypto/bcrypt"
	"scoi-web/internal/repository"
)

type authService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *authService {
	return &authService{repo}
}

func (a *authService) SignIn(input SignInInput) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	id, err := a.repo.CreateUser(input.Email, string(hashedPassword))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *authService) SignUp(input SignInInput) (int, error) {
	id, err := a.repo.GetUser(input.Email, input.Password)
	if err != nil {
		return 0, err
	}
	return id, nil
}
