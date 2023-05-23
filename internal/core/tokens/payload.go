package tokens

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TokenPayload struct {
	TokenID   uuid.UUID `json:"token_id"`   // Token id
	UserID    uuid.UUID `json:"client_id"`  // И так понятно че комент на каждую строку писать?
	RtokenID  uuid.UUID `json:"rtoken_id"`  // Refresh Token id
	IssuedAt  time.Time `json:"issued_at"`  // Created at
	ExpiresAt time.Time `json:"expires_at"` // Дата когда токен должен просрочиться
}

func NewPayload(userID uuid.UUID, tLifeTime time.Duration, rTokenID uuid.UUID) (*TokenPayload, error) {
	tokenId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	newPayload := TokenPayload{
		TokenID:   tokenId,
		UserID:    userID,
		RtokenID:  rTokenID,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(tLifeTime),
	}

	return &newPayload, nil
}

func (p *TokenPayload) Valid() error {
	if time.Now().After(p.ExpiresAt) {
		return errors.New("token expired")
	}
	return nil
}
