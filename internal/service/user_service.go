package service

import "scoi-web/internal/repository"

type UserService struct {
	repos *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repos: repo}
}

func (u *UserService) CreateUser() {}
func (u *UserService) UpdateUser() {}
