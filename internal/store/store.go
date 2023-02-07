package store

import (
	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
)

type Store struct {
	Storage models.Users
}

func NewStore() *Store {
	s := &Store{
		Storage: models.Users{},
	}

	return s
}

func (s *Store) AddUser(u models.User) {
	s.Storage.Users = append(s.Storage.Users, u)
}

func (s *Store) GetUsers() models.Users {

	return s.Storage
}
