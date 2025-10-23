package main

import (
	"log"

	"github.com/finance/apiService/config"
	"github.com/finance/apiService/internal/app"
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
