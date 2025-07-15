package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds application configuration.
type Config struct {
	Rights      string
	Timeout     time.Duration
	GoAPIPath   string
	FrontendURL string
}

// LoadConfig loads configuration from environment.
func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		Rights: os.Getenv("RIGHTS"),
		Timeout: func() time.Duration {
			timeout := os.Getenv("TIMEOUT")
			intTimeout, _ := strconv.Atoi(timeout)
			return time.Duration(intTimeout) * time.Second
		}(),
		GoAPIPath:   os.Getenv("GO_API_URL"),
		FrontendURL: os.Getenv("FRONTEND_URL"),
	}
}
