package postgres

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser() {

}

func (u *UserRepository) UpdateUser() {

}
