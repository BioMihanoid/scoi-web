package service

import "scoi-web/internal/repository"

type UserServiceInterface interface {
	CreateUser()
	UpdateUser()
}

type Service struct {
	UserServiceInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserServiceInterface: NewUserService(repo),
	}
}
