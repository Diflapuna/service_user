package store

import (
	"time"

	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
	"github.com/google/uuid"
)

// кривая хуйня надо будет починить рефактором
const (
	RtLifetime = time.Hour * 10
)

func (s *Store) LoginUser(email string, password string) (uuid.UUID, error) {
	user := &models.User{}

	row := s.DB.QueryRowx(
		"SELECT id FROM users WHERE email = $1 AND password = $2;",
		email, password,
	)
	if err := row.StructScan(user); err != nil {
		s.Logger.Errorf("Can't scan user struct %w", err)
		return uuid.Nil, err
	}

	if user.Id == uuid.Nil {
		s.Logger.Error("Invalid login/password")
	}

	s.Logger.Info("Welcome to the club, buddy!")
	return user.Id, nil
}

// надо будет это расхардкодить и положить сюда функцию парса токенов пожалуйста
func (s *Store) CreateRefreshSession(rToken string, rTokenID uuid.UUID, userID uuid.UUID) error {
	rCreatedAt := time.Now()
	lifetime := rCreatedAt.Add(RtLifetime)
	s.DB.QueryRowx(
		"INSERT INTO refresh_sessions (id_client, id_refresh_token, issued_at, expires_in) VALUES ($1, $2, $3, $4);",
		userID, rTokenID, rCreatedAt, lifetime,
	)

	return nil
}
