package store

import (
	"database/sql"

	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Store struct {
	Storage models.Users
	DB      *sqlx.DB
	Logger  *zap.SugaredLogger
}

func NewStore(l *zap.SugaredLogger) *Store {
	s := &Store{
		Storage: models.Users{},
		DB:      newDB(l),
		Logger:  l,
	}
	return s
}

func newDB(l *zap.SugaredLogger) *sqlx.DB {
	urlExample := "postgres://user:123456789@localhost:5430/test1"
	sql.Register("wrapper", stdlib.GetDefaultDriver())
	wdb, err := sql.Open("wrapper", urlExample)
	if err != nil {
		l.Fatalf("Can't connect to DB %w", err)
	}
	db := sqlx.NewDb(wdb, "wrapper")

	return db
}
