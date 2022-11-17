package handler

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed redoc.html
var html string

func NewReDoc() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, html)
	}
}
