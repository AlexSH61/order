package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/AlexSH61/order/internal/logger"
	"github.com/kelseyhightower/envconfig"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

type config struct {
	Env            string
	Goose_Driver   string `envconfig:GOOSE_DRIVER`
	Goose_DBstring string `envconfig:GOOSE_DBSTRING`
}

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		fmt.Println("Error creating logger:", err)
		os.Exit(1)
	}
	defer log.Sync()

	if err = run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return err
	}
	log.Infow("startup", "STATUS!", "OK")
	log.Infow("cfg", "ENV", cfg.Env)

	db, err := sql.Open(cfg.Goose_Driver, cfg.Goose_DBstring)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return err
	}
	if err := goose.Up(db, "internal/migrations/"); err != nil {
		return err
	}
	return nil
}
