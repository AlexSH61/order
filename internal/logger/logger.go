package logger

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	log, err := config.Build()
	if err != nil {
		return nil, err
	}
	return log.Sugar(), nil
}
