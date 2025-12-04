package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type (
	// Config -.
	Config struct {
		App             App
		HTTP            HTTP
		Log             Log
		PG              PG
		RMQ             RMQ
		RMQReceivers    RMQReceivers
		WebapiReceivers WebapiReceivers
		Metrics         Metrics
		Swagger         Swagger
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port          string `env:"HTTP_PORT,required"`
		MaxUploadSize int64  `env:"HTTP_MAX_UPLOAD_SIZE" envDefault:"15728640"` // Default 15 MB
	}

	// Log -.
	Log struct {
		Level    string `env:"LOG_LEVEL,required"`
		FileName string `env:"LOG_FILE_NAME,required"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env:"RMQ_RPC_SERVER_EXCHANGE,required"`
		ClientExchange string `env:"RMQ_RPC_CLIENT_EXCHANGE,required"`
		URL            string `env:"RMQ_URL,required"`
		ClientPrefix   string `env:"RMQ_RPC_CLIENT_PREFIX,required"`
	}

	// RMQReceivers -.
	RMQReceivers struct {
		BackService string `env:"RMQ_RECEIVERS_BACKEND_SERVICE,required"`
	}

	// WebapiReceivers -.
	WebapiReceivers struct {
		BackService string `env:"WEBAPI_RECEIVERS_BACKEND_SERVICE,required"`
	}

	// Metrics -.
	Metrics struct {
		Enabled bool `env:"METRICS_ENABLED" envDefault:"true"`
	}

	// Swagger -.
	Swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Could not loading .env file")
	}

	// Parse environment variables into structs
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
