package app

import (
	"shared/app"
)

type Config struct {
	AppMode   app.Mode
	SentryDSN string
	ApiUrl    string
	SignKey   string
	DbPath    string
}

func NewConfig() *Config {
	return &Config{
		app.Mode(app.MustEnv("APP_MODE")),
		app.MustEnv("SENTRY_DSN"),
		app.MustEnv("API_URL"),
		app.MustEnv("API_SIGN_KEY"),
		app.MustEnv("DB_PATH"),
	}
}
