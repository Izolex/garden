package http

import (
	"main/http/handler"
	"main/work"
	"net/http"
	sharedApi "shared/api"
	"shared/api/middleware"
	"shared/api/sign"
	"shared/app/logger"
)

func NewHandler(manager work.Manager, logger logger.Logger, signService sign.Service) http.Handler {
	engine := sharedApi.NewEngine(logger)

	engine.Use(middleware.NewSignature(signService))

	engine.POST("/work", work.NewWorkPOST(manager))
	engine.GET("/status", handler.NewStatusGET(manager))

	return engine
}
