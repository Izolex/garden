package handler

import "github.com/gin-gonic/gin"

func NewError() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic("Palačinky se salámem!!!")
	}
}
