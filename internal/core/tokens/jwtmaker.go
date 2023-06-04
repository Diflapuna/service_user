package tokens

import (
	"fmt"
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

func (j *JWTMaker) ParseToken(token string) (*TokenPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("wrong signing method")
		}

		return []byte(j.secretKey), nil
	}
	parsedToken, err := jwt.ParseWithClaims(token, &TokenPayload{}, keyFunc)
	if err != nil {
		verr, _ := err.(*jwt.ValidationError) //nolint: errorlint
		if verr.Errors == jwt.ValidationErrorClaimsInvalid {
			_, ok := parsedToken.Claims.(*TokenPayload)
			if !ok {
				return nil, fmt.Errorf("token expired: %v", err)
			}
		}

		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	payload, ok := parsedToken.Claims.(*TokenPayload)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	return payload, nil
}

func ParseToken(token string) (*TokenPayload, error) {
	maker := NewJWTMaker("iDidNotHitHer")

	payload, err := maker.ParseToken(token)
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	return payload, nil
}

func (j *JWTMaker) ParseExpiredToken(token string) (*TokenPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("wrong signing method")
		}

		return []byte(j.secretKey), nil
	}

	parsedToken, _ := jwt.ParseWithClaims(token, &TokenPayload{}, keyFunc)

	payload, ok := parsedToken.Claims.(*TokenPayload)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	return payload, nil

}
