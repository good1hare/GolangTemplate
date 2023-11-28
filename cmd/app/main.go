package main

import (
	"github.com/good1hare/GolangTemplate/configs"
	"github.com/good1hare/GolangTemplate/internal/app"
	"log"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
