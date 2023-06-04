package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
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

	log := NewLogger()

	s := &Service{
		Log:    log,
		Store:  store.NewStore(log),
		router: mux.NewRouter(),
	}
	s.registerHandlers()
	s.GreetGateway()
	return s
}

func (s *Service) Start() error {
	s.Log.Info("Starting listening on port 6969")
	return fmt.Errorf("failed to start service %w", http.ListenAndServe(":6969", s.router))
}

// gonna remove hardcoding in the future, this is not okay but i don't have time now
func (s *Service) GreetGateway() {
	greet := &models.Greeting{}
	s.Log.Info("Trying to connect to gateway...")
	resp, err := http.Get("http://localhost:1337/hello")
	if err != nil {
		s.Log.Fatal("Could not connect to API gateway: ", err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(greet)
	if err != nil {
		s.Log.Fatal("Could not decode response from GW: ", err)
	}
	if greet.Greeting != "O hi Mark!" {
		s.Log.Fatal("Wrong greeting, expected 'O hi Mark!', got ", greet.Greeting)
	} else {
		s.Log.Info("Sending handlers list to API gateway...")
		s.sendHandlersList()
	}
}

func (s *Service) sendHandlersList() {
	handlers, err := os.Open("handlers/handlers.json")
	if err != nil {
		s.Log.Fatal("Could not open handler list", err)
	}
	reader, err := io.ReadAll(handlers)
	if err != nil {
		s.Log.Fatal("Could not open handler list", err)
	}

	resp, err := http.Post("http://localhost:1337/handlers", "application/json", bytes.NewBuffer(reader))
	if err != nil {
		s.Log.Fatal("Could not send handler list", err)
	}
	if resp.StatusCode != http.StatusOK {
		s.Log.Fatal("Because fuck you that's why: ", err)
	}
	s.Log.Info("Handlers sent successfully")
}
