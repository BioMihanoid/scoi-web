package repository

import (
	"github.com/jmoiron/sqlx"
	"scoi-web/internal/repository/postgres"
)

type UserRepositoryInterface interface {
	CreateUser()
	UpdateUser()
}

type Repository struct {
	UserRepositoryInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepositoryInterface: postgres.NewUserRepository(db),
	}
}
