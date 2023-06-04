package tokens

import (
	"log"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	key := []byte(j.secretKey)
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Print(err)
		return "", uuid.Nil, err
	}

	return tokenString, payload.RtokenID, nil
}
