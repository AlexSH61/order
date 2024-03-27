package pkg

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Env string
}

func InitLogger() (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()

	log, err := config.Build()
	if err != nil {
		return nil, err
	}
	return log.Sugar(), nil
}

func Run(log *zap.SugaredLogger) error {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return err
	}

	log.Infow("startup", "STATUS", "OK!")
	log.Infow("cfg", "ENV", cfg.Env)
	return nil
}
