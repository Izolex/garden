package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")
		if contentType != "application/json" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Next()
	}
}
