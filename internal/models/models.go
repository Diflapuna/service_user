package models

import (
	"time"

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

type Tokens struct {
	Acsess  string `json:"acsess"`
	Refresh string `json:"refresh"`
}

type RefreshSession struct {
	ID             int       `db:"id"`
	ClientID       uuid.UUID `db:"id_client"`
	RefreshTokenID uuid.UUID `db:"id_refresh_token"`
	IssuedAt       time.Time `db:"issued_at"`
	ExpiresIn      time.Time `db:"expires_in"`
}
