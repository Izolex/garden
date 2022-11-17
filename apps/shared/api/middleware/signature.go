package middleware

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/api/sign"
)

func NewSignature(service sign.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("x-signature")

		signStr, err := hex.DecodeString(header)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !service.Verify(c.Request, signStr) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
