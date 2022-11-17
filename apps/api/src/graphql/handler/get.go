package handler

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed graphiql.html
var html string

func NewGET() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, html)
	}
}
