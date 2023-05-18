package store

import (
	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
	"github.com/google/uuid"
)

func (s *Store) GetUsers() (models.Users, error) {
	users := models.Users{Users: []models.User{}}
	rows, err := s.DB.Queryx(
		"SELECT * FROM users;",
	)
	if err != nil {
		s.Logger.Errorf("Can't get users %w", err)
		return models.Users{}, err
	}
	user := &models.User{}
	for rows.Next() {
		rows.StructScan(user)
		users.Users = append(users.Users, *user)
	}

	return users, nil
}

func (s *Store) AddUser(u models.User) {
	s.Storage.Users = append(s.Storage.Users, u)

	s.DB.QueryRow(
		"INSERT INTO users (id, name, email, password, about) VALUES ($1, $2, $3, $4, $5);",
		u.Id, u.Name, u.Email, u.Password, u.About,
	)
}

func (s *Store) DeleteUser() {

}

func (s *Store) EditPassword(newPassword string, u uuid.UUID) error {
	_, err := s.DB.Exec(
		"UPDATE users SET password = $1 WHERE id = $2;",
		newPassword, u,
	)

	if err != nil {
		s.Logger.Errorf("Can't change password %w", err)
		return err
	}

	return nil
}

func (s *Store) EditEmail(newEmail string, u uuid.UUID) error {
	_, err := s.DB.Exec(
		"UPDATE users SET email = $1 WHERE id = $2;",
		newEmail, u,
	)

	if err != nil {
		s.Logger.Errorf("Can't change email %w", err)
		return err
	}

	return nil
}

func (s *Store) EditAbout(newAbout string, u uuid.UUID) error {
	_, err := s.DB.Exec(
		"UPDATE users SET about = $1 WHERE id = $2;",
		newAbout, u,
	)

	if err != nil {
		s.Logger.Errorf("Can't change about info %w", err)
		return err
	}

	return nil
}
