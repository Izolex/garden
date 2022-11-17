package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/api/handler"
	"shared/api/middleware"
	"shared/app/logger"
)

func NewEngine(logger logger.Logger) *gin.Engine {
	engine := gin.New()

	engine.Use(middleware.NewRecovery(logger))
	engine.Use(middleware.NewAccessControl())

	engine.NoMethod(func(c *gin.Context) {
		c.Status(http.StatusMethodNotAllowed)
	})
	engine.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})

	engine.GET("/make-error", handler.NewError())

	return engine
}
