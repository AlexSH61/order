package main

import (
	"fmt"
	"os"

	"github.com/AlexSH61/1project/pkg"
)

type Config struct {
	Env string
}

func main() {
	log, err := pkg.InitLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := pkg.Run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		os.Exit(1)
	}
}
