package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/app/logger"
)

func NewRecovery(loggerIns logger.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		defer func() {
			c.AbortWithStatus(http.StatusInternalServerError)
			ctx := context.WithValue(context.Background(), logger.ErrorCtxHttpRequest, c.Request)
			loggerIns.ErrorCtx(fmt.Errorf("%v", recovered), ctx)
		}()

		c.Next()
	})
}
