package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
	"github.com/google/uuid"
)

func (s *Service) registerHandlers() {
	s.router.HandleFunc("/users", s.RegisterUser()).Methods("POST")
	s.router.HandleFunc("/users", s.GetAllUsers()).Methods("GET")
}

func (s *Service) GetAllUsers() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.Store.GetUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(users); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *Service) RegisterUser() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		user.Id = uuid.New()
		s.Store.AddUser(*user)
		w.WriteHeader(http.StatusCreated)
		fmt.Println(s.Store.Storage.Users)
	}
}
