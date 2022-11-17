//todo go : generate swag init --output ./docs/generated && swag fmt
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main/app"
	"main/graphql/schema"
	appHttp "main/http"
	sharedApi "shared/api"
	"shared/api/sign"
	sharedApp "shared/app"
	"shared/measurement"
	"shared/model/database"
)

// @Title        Mung bean sprout garden REST API :))))
// @Version      1.0
// @Description  Mmmmm it tastes so good..

// @contact.name   Jan Tuzil
// @contact.email  info@jantuzil.cz
// @contact.url    https://jantuzil.cz
func main() {
	fmt.Println("Hi!")

	config := app.NewConfig()

	logger := sharedApp.NewLogger(config.AppMode, config.SentryDSN)
	defer logger.Stop()

	db, sqlDB := database.NewMySQL(config.MysqlDSN)
	defer sqlDB.Close()

	schemaProvider := schema.NewProvider(db)
	signService := sign.NewService(config.SignKey)
	measurementStorageService := measurement.NewStorage(db)

	handler := appHttp.NewHandler(sqlDB, logger, schemaProvider, signService, measurementStorageService)
	httpServer := sharedApi.NewServer(":80", handler, logger)
	go httpServer.Run()

	<-sharedApp.OnInterrupt()

	fmt.Println("Gracefully shutdown...")
	httpServer.Shutdown()
}
