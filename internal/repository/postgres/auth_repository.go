package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"scoi-web/internal/models"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) CreateUser(email string, password string) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s(email, hashed_password) VALUES ($1, $2) returning id", usersTable)
	row := a.db.QueryRow(query, email, password)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthRepository) GetUser(email string, password string) (int, error) {
	var user models.User
	query := fmt.Sprintf("select * from %s where email = $1", usersTable)

	rows, err := a.db.Query(query, email)
	if err != nil {
		return user.Id, err
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.SecondName, &user.Email, &user.Password, &user.CreatedAt)
	}
	if user.Id == 0 {
		return user.Id, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, fmt.Errorf("invalid password")
	}
	return user.Id, nil
}
