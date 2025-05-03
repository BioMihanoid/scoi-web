package repository

import (
	"github.com/jmoiron/sqlx"
	"scoi-web/internal/repository/postgres"
)

type AuthRepositoryInterface interface {
	CreateUser(email string, password string) (int, error)
	GetUser(email string, password string) (int, error)
}

type Repository struct {
	AuthRepositoryInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthRepositoryInterface: postgres.NewAuthRepository(db),
	}
}
