package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
}

type Users struct {
	Users []User `json:"users"`
}
