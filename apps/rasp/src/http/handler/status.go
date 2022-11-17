package handler

import (
	"github.com/gin-gonic/gin"
	"main/work"
	"net/http"
)

type StatusResponse struct {
	RunnersCount int `json:"runners_count"`
}

func NewStatusGET(manager work.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, StatusResponse{
			RunnersCount: manager.GetRunnersCount(),
		})
	}
}
