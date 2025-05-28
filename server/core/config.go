package core

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	JwtSecret   string
}

func MustGetConfig() *Config {
	// Production will cause godotenv to fail so just log as warning.
	if err := godotenv.Load(); err != nil {
		Errorf(WrappedError(err, "failed to load config"))
	}
	return &Config{
		DatabaseUrl: os.Getenv("DB_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}
