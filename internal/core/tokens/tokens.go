package tokens

import (
	"time"

	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
	"github.com/google/uuid"
)

const (
	secretKey  = "iDidNotHitHerIDidNaaat"
	tLifetime  = time.Minute * 10
	rtLifetime = time.Hour * 10
)

func GenerateTokenPair(userID uuid.UUID) (*models.Tokens, error) {
	maker := NewJWTMaker(secretKey)

	rToken, rTokenID, err := maker.CreateToken(userID, uuid.Nil, rtLifetime)
	if err != nil {
		return nil, err
	}
	aToken, _, err := maker.CreateToken(userID, rTokenID, tLifetime)
	if err != nil {
		return nil, err
	}

	return &models.Tokens{Acsess: aToken, Refresh: rToken}, nil
}
