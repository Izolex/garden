//go:generate mockgen -source=logger.go -destination=mock/logger.go -package=mock
package logger

import "context"

type ErrorCtx int

const ErrorCtxHttpRequest ErrorCtx = iota

type Logger interface {
	Error(error)
	ErrorCtx(err error, ctx context.Context)
	Info(string)
}

type Stopper interface {
	Stop()
}
