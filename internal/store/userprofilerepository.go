package store

import (
	"log"

	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
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
