package app

import (
	sentry "github.com/getsentry/sentry-go"
	"os"
	"shared/app/logger"
)

func NewLogger(appMode Mode, sentryDsn string) interface {
	logger.Logger
	logger.Stopper
} {
	var loggers []logger.Logger
	loggers = append(loggers, logger.NewFileLogger(os.Stdout))
	if appMode == ModeProduction || appMode == ModeStaging {
		sentryLogger := logger.NewSentryLogger(sentry.ClientOptions{
			Dsn:         sentryDsn,
			Environment: string(appMode),
			Debug:       appMode == ModeStaging,
		})
		loggers = append(loggers, sentryLogger)
	}
	return logger.NewMultiLogger(loggers)
}
