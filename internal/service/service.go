package service

import (
	"fmt"
	"net/http"

	"github.com/NotYourAverageFuckingMisery/animello/internal/store"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Service struct {
	Log    *zap.SugaredLogger
	Store  *store.Store
	router *mux.Router
}

func NewService() *Service {
	s := &Service{
		Log:    NewLogger(),
		Store:  store.NewStore(),
		router: mux.NewRouter(),
	}
	s.registerHandlers()
	return s
}

func (s *Service) Start() error {
	s.Log.Info("Starting listening on port 6969")
	return fmt.Errorf("failed to start service %w", http.ListenAndServe(":6969", s.router))
}
