package http

import (
	"database/sql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "main/docs/generated"
	docsHandler "main/docs/handler"
	graphqlHandler "main/graphql/handler"
	"main/graphql/schema"
	"main/http/handler"
	measurementHandler "main/modules/measurement/handler"
	"net/http"
	sharedApi "shared/api"
	"shared/api/middleware"
	"shared/api/sign"
	"shared/app/logger"
	"shared/measurement"
)

func NewHandler(
	db *sql.DB,
	logger logger.Logger,
	schemaProvider *schema.Provider,
	signService sign.Service,
	measurementStorageService measurement.Storage,
) http.Handler {
	signatureMiddleware := middleware.NewSignature(signService)
	jsonMiddleware := middleware.NewJson()

	engine := sharedApi.NewEngine(logger)
	engine.GET("/status", handler.NewStatus(db))

	api := engine.Group("/api")
	{
		api.Group("/v1/measurement").POST("", jsonMiddleware, signatureMiddleware, measurementHandler.NewPOST(measurementStorageService))
	}
	doc := engine.Group("/doc")
	{
		doc.GET("/redoc", docsHandler.NewReDoc())
		doc.GET("/swag/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	graphQL := engine.Group("/graphql")
	{
		graphQL.GET("", graphqlHandler.NewGET())
		graphQL.POST("", graphqlHandler.NewPOST(schemaProvider))
	}

	return engine
}
