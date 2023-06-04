package tokens

import (
	"time"

	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
	"github.com/NotYourAverageFuckingMisery/animello/internal/store"
	"github.com/google/uuid"
)

const (
	TLifetime  = time.Minute * 10
	RtLifetime = time.Hour * 10
)

func GenerateTokenPair(userID uuid.UUID) (*models.Tokens, uuid.UUID, error) {
	secretKey := "iDidNotHitHer"
	maker := NewJWTMaker(secretKey)

	rToken, rTokenID, err := maker.CreateToken(userID, uuid.Nil, RtLifetime)
	if err != nil {
		return nil, uuid.Nil, err
	}
	aToken, _, err := maker.CreateToken(userID, rTokenID, TLifetime)
	if err != nil {
		return nil, uuid.Nil, err
	}

	return &models.Tokens{Acsess: aToken, Refresh: rToken}, rTokenID, nil
}

func Refresh(rToken string, s *store.Store) (*models.Tokens, error) {
	maker := NewJWTMaker("iDidNotHitHer")
	payload, err := maker.ParseExpiredToken(rToken)
	if err != nil {
		s.Logger.Error("Failed to parse expired token lol")
		return nil, err
	}

	aToken, _, err := maker.CreateToken(payload.UserID, payload.RtokenID, TLifetime)
	if err != nil {
		return nil, err
	}

	return &models.Tokens{Acsess: aToken}, nil
}
