package store

import (
	"database/sql"
	"log"

	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	Storage models.Users
	DB      *sqlx.DB
}

func NewStore() *Store {
	s := &Store{
		Storage: models.Users{},
		DB:      newDB(),
	}

	return s
}

func (s *Store) AddUser(u models.User) {
	s.Storage.Users = append(s.Storage.Users, u)
}

func (s *Store) GetUsers() models.Users {

	return s.Storage
}

func newDB() *sqlx.DB {
	urlExample := "postgres://postgres:14881488@localhost:5432/frog"
	sql.Register("wrapper", stdlib.GetDefaultDriver())
	wdb, err := sql.Open("wrapper", urlExample)
	if err != nil {
		log.Fatal(err)
	}
	db := sqlx.NewDb(wdb, "wrapper")

	return db
}
