package service

import (
	"scoi-web/internal/repository"
)

type SignInInput struct {
	Email    string
	Password string
}

type UserServiceInterface interface {
}

type AuthServiceInterface interface {
	SignIn(input SignInInput) (int, error)
	SignUp(input SignInInput) (int, error)
}

type Service struct {
	TokenManager
	UserServiceInterface
	AuthServiceInterface
}

func NewService(repo *repository.Repository, jwtKey string) *Service {
	return &Service{
		TokenManager:         NewManager(jwtKey),
		UserServiceInterface: NewUserService(repo),
		AuthServiceInterface: NewAuthService(repo),
	}
}
