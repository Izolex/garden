package api

import (
	"context"
	"net/http"
	"shared/app/logger"
	"time"
)

type Server struct {
	logger     logger.Logger
	httpServer *http.Server
}

func NewServer(addr string, handler http.Handler, logger logger.Logger) *Server {
	return &Server{
		logger: logger,
		httpServer: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Run() {
	if err := s.httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error(err)
	}
}
