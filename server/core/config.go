package core

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerMode int

const (
	SERVER_MODE_DEV ServerMode = iota
	SERVER_MODE_PROD
)

type Config struct {
	DatabaseUrl   string
	JwtSecret     string
	ServerMode    ServerMode
	SpotifyConfig *SpotifyConfig
	YoutubeConfig *YoutubeConfig
}

type SpotifyConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

type YoutubeConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

func MustGetConfig() *Config {
	// Production will cause godotenv to fail so just log as warning.
	if err := godotenv.Load(); err != nil {
		Errorf(WrappedError(err, "failed to load config"))
	}
	var serverMode ServerMode
	serverModeStr := os.Getenv("SERVER_MODE")
	switch serverModeStr {
	case "dev":
		serverMode = SERVER_MODE_DEV
	case "prod":
		serverMode = SERVER_MODE_PROD
	default:
		// Safer to fallback to production.
		serverMode = SERVER_MODE_PROD
	}
	return &Config{
		DatabaseUrl: os.Getenv("DB_URL"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
		ServerMode:  serverMode,
		SpotifyConfig: &SpotifyConfig{
			ClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			RedirectUri:  os.Getenv("SPOTIFY_REDIRECT_URI"),
		},
		YoutubeConfig: &YoutubeConfig{
			ClientId:     os.Getenv("YOUTUBE_CLIENT_ID"),
			ClientSecret: os.Getenv("YOUTUBE_CLIENT_SECRET"),
			RedirectUri:  os.Getenv("YOUTUBE_REDIRECT_URI"),
		},
	}
}
