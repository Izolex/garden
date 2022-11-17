package logger

import (
	"context"
	"log"
	"os"
)

type fileLogger struct {
	logger *log.Logger
}

func (l *fileLogger) Error(err error) {
	l.logger.Println(err)
}

func (l *fileLogger) ErrorCtx(err error, ctx context.Context) {
	l.logger.Printf("%v - CONTEXT: - %v\n", err, ctx)
}

func (l *fileLogger) Info(msg string) {
	l.logger.Println(msg)
}

func NewFileLogger(file *os.File) *fileLogger {
	return &fileLogger{log.New(file, "", log.LstdFlags)}
}
