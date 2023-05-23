package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	About    string    `json:"about" db:"about"`
}

type Users struct {
	Users []User `json:"users"`
}

type Greeting struct {
	Greeting string `json:"greeting"`
}
