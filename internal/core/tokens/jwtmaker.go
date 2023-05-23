package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{secretKey: secretKey}
}

func (j *JWTMaker) CreateToken(userID uuid.UUID, rTokenID uuid.UUID, lifetime time.Duration) (string, uuid.UUID, error) {
	payload, err := NewPayload(userID, lifetime, rTokenID)
	if err != nil {
		return "", uuid.Nil, err
	}
	// нахуй это надо я не понимаю может потом уберу если оно не нужно
	if rTokenID == uuid.Nil {
		payload.RtokenID = payload.TokenID
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, payload)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", uuid.Nil, err
	}

	return tokenString, payload.RtokenID, nil
}
