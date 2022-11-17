package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go/relay"
	"main/graphql/schema"
)

func NewPOST(schemaProvider *schema.Provider) gin.HandlerFunc {
	return gin.WrapH(&relay.Handler{
		Schema: schemaProvider.GetSchema(),
	})
}
