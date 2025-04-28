package models

import "time"

type User struct {
	ID         int64     `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
}
