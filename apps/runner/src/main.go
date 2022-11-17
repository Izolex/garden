package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"main/app"
	appHttp "main/http"
	"main/plan"
	"main/planValue"
	"main/raspberry"
	"main/work"
	"net/http"
	sharedApi "shared/api"
	"shared/api/request"
	"shared/api/sign"
	sharedApp "shared/app"
	"shared/model/database"
	planModel "shared/model/entity/plan"
	sharedWork "shared/work"
)

func main() {
	fmt.Println("Hi!")

	config := app.NewConfig()

	logger := sharedApp.NewLogger(config.AppMode, config.SentryDSN)
	defer logger.Stop()

	db, sqlDB := database.NewMySQL(config.MysqlDSN)
	defer sqlDB.Close()

	signService := sign.NewService(config.SignKey)
	requestFactory := request.NewFactory(signService)
	httpClient := &http.Client{}
	workStorage := sharedWork.NewStorage(db)
	raspService := raspberry.NewService(httpClient, logger, workStorage, requestFactory)
	planValueModel := planValue.NewModel(db)

	plans := newPlanMap(db, raspService, planValueModel)
	planner := work.NewPlanner(config.TickerDuration, db, logger, plans)
	go planner.Run()

	handler := appHttp.NewHandler(sqlDB, logger)
	httpServer := sharedApi.NewServer(":80", handler, logger)
	go httpServer.Run()

	<-sharedApp.OnInterrupt()

	fmt.Println("Shutdown...")
}

func newPlanMap(db *gorm.DB, raspService raspberry.Service, planValueModel planValue.Model) work.PlanMap {
	return work.PlanMap{
		planModel.MungBeanSprouts: plan.NewMungBeanSprouts(db, raspService, planValueModel),
	}
}
