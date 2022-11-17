package app

import (
	"shared/app"
)

type Config struct {
	AppMode        app.Mode
	MysqlDSN       string
	SentryDSN      string
	ApiUrl         string
	TickerDuration int
	SignKey        string
}

func NewConfig() *Config {
	return &Config{
		app.Mode(app.MustEnv("APP_MODE")),
		app.MustEnv("MYSQL_DSN"),
		app.MustEnv("SENTRY_DSN"),
		app.MustEnv("API_URL"),
		app.MustEnvInt("TICKER_DURATION"), // ms
		app.MustEnv("API_SIGN_KEY"),
	}
}
