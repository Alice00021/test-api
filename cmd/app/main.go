package main

import (
	"log"

	"github.com/Alice00021/test_api/config"
	"github.com/Alice00021/test_api/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
