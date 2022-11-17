package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewStatus(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"database": db.Ping() == nil,
		})
	}
}
