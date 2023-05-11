package store

import (
	"log"

	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
	"github.com/google/uuid"
)

func (s *Store) GetUsers() (models.Users, error) {
	users := models.Users{Users: []models.User{}}
	rows, err := s.DB.Queryx(
		"SELECT * FROM users;",
	)
	if err != nil {
		log.Println(err)
		return models.Users{}, err
	}
	user := models.User{}
	for rows.Next() {
		rows.StructScan(user)
		users.Users = append(users.Users, user)
	}

	return users, nil
}

func (s *Store) AddUser(u models.User) {
	s.Storage.Users = append(s.Storage.Users, u)

	s.DB.QueryRow(
		"INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4);",
		u.Id, u.Name, u.Email, u.Password,
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
