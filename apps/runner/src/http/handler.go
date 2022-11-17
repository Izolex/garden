package http

import (
	"database/sql"
	"main/http/handler"
	"net/http"
	sharedApi "shared/api"
	"shared/app/logger"
)

func NewHandler(db *sql.DB, logger logger.Logger) http.Handler {
	engine := sharedApi.NewEngine(logger)
	engine.GET("/status", handler.NewStatus(db))

	return engine
}
