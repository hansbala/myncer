package core

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
}

func MustGetConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(WrappedError(err, "failed to load config"))
	}
	return &Config{
		DatabaseUrl: os.Getenv("DB_URL"),
	}
}
