package logger

import "context"

type multiLogger struct {
	list []Logger
}

func (ml *multiLogger) Error(err error) {
	for _, logger := range ml.list {
		logger.Error(err)
	}
}

func (ml *multiLogger) ErrorCtx(err error, ctx context.Context) {
	for _, logger := range ml.list {
		logger.ErrorCtx(err, ctx)
	}
}

func (ml *multiLogger) Info(info string) {
	for _, logger := range ml.list {
		logger.Info(info)
	}
}

func (ml *multiLogger) Stop() {
	for _, logger := range ml.list {
		if stopper, ok := logger.(Stopper); ok {
			stopper.Stop()
		}
	}
}

func NewMultiLogger(list []Logger) *multiLogger {
	return &multiLogger{list}
}
