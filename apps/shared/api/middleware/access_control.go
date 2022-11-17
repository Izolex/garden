package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewAccessControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		if c.Request.Method == http.MethodOptions {
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "content-type,x-signature,sentry-trace")
			c.Header("Access-Control-Max-Age", "86400")

			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
