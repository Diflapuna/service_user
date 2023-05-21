package service

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {

	logger, err := zap.NewDevelopment()

	if err != nil {
		logger.Fatal("Can't create logger")
	}
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("Logger is started and ready to log.")

	return sugar
}
