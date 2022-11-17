package logger

import (
	"context"
	"fmt"
	"github.com/certifi/gocertifi"
	"github.com/getsentry/sentry-go"
	"net/http"
	"time"
)

type sentryLogger struct {
	hub *sentry.Hub
}

func (sl *sentryLogger) Error(err error) {
	sl.hub.Clone().CaptureException(err)
}

func (sl *sentryLogger) ErrorCtx(err error, ctx context.Context) {
	hub := sl.hub.Clone()
	if req, ok := ctx.Value(ErrorCtxHttpRequest).(*http.Request); ok {
		hub.Scope().SetRequest(req)
	}
	hub.RecoverWithContext(ctx, err)
}

func (sl *sentryLogger) Info(msg string) {
	sl.hub.Clone().CaptureMessage(msg)
}

func (sl *sentryLogger) Stop() {
	if !sl.hub.Flush(2 * time.Second) {
		fmt.Println("sentry logger: stopper flush not successfull")
	}
}

func NewSentryLogger(options sentry.ClientOptions) *sentryLogger {
	rootCAs, err := gocertifi.CACerts()
	if err != nil {
		panic(fmt.Errorf("sentry logger: could not load CA certificates: %v\n", err))
	} else {
		options.CaCerts = rootCAs
	}

	client, err := sentry.NewClient(options)
	if err != nil {
		panic(err)
	}
	scope := sentry.NewScope()
	hub := sentry.NewHub(client, scope)
	return &sentryLogger{hub}
}
