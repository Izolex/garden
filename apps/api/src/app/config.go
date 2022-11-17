package app

import (
	"shared/app"
)

type Config struct {
	AppMode   app.Mode
	SentryDSN string
	MysqlDSN  string
	SignKey   string
}

func NewConfig() *Config {
	return &Config{
		app.Mode(app.MustEnv("APP_MODE")),
		app.MustEnv("SENTRY_DSN"),
		app.MustEnv("MYSQL_DSN"),
		app.MustEnv("API_SIGN_KEY"),
	}
}
