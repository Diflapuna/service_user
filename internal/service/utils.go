package service

import (
	"go.uber.org/zap"
)

/*
func (s *Service) error(w http.ResponseWriter, code int, err error, ctx context.Context) {

}
*/
func NewLogger() *zap.SugaredLogger {

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()

	return sugar
}
