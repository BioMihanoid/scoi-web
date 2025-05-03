package models

import "time"

type User struct {
	Id         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Email      string    `json:"email"`
	Password   string    `json:"hashed_password"`
	CreatedAt  time.Time `json:"created_at"`
}
