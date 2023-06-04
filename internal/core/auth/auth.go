package auth

import (
	"github.com/NotYourAverageFuckingMisery/animello/internal/core/tokens"
	"github.com/NotYourAverageFuckingMisery/animello/internal/models"
	"github.com/NotYourAverageFuckingMisery/animello/internal/store"
)

func Auth(s *store.Store, email string, pwd string) (*models.Tokens, error) {
	//bytesPwd := []byte(pwd)

	userId, err := s.LoginUser(email, pwd)
	if err != nil {
		return nil, err
	}
	tokenPair, refreshId, err := tokens.GenerateTokenPair(userId)
	if err != nil {
		return nil, err
	}
	s.CreateRefreshSession(tokenPair.Refresh, refreshId, userId)
	return tokenPair, nil
}
