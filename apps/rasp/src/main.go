package main

import (
	"fmt"
	"main/api"
	"main/app"
	appHttp "main/http"
	"main/io"
	"main/job"
	"main/work"
	"net/http"
	"periph.io/x/host/v3"
	sharedApi "shared/api"
	"shared/api/request"
	"shared/api/sign"
	sharedApp "shared/app"
	"shared/model/database"
	jobModel "shared/model/entity/job"
	"time"
)

func main() {
	fmt.Println("Hi!")

	_, err := host.Init()
	if err != nil {
		panic(err)
	}

	config := app.NewConfig()
	logger := sharedApp.NewLogger(config.AppMode, config.SentryDSN)
	pins := io.NewPins(config.AppMode, logger)

	db, sqlDB := database.NewSQLite(config.DbPath)
	defer sqlDB.Close()
	requestModel := api.NewRequestModel(db)

	httpClient := &http.Client{}
	signService := sign.NewService(config.SignKey)
	requestFactory := request.NewFactory(signService)

	apiStore := api.NewStore(requestModel)
	apiSender := api.NewSender(config.ApiUrl, 5*time.Second, httpClient, requestFactory, requestModel, logger)
	go apiSender.Run()

	jobMap := newJobMap(apiStore)
	jobsManager := work.NewManager(pins, jobMap)
	go jobsManager.Run()
	defer jobsManager.StopAll()

	handler := appHttp.NewHandler(jobsManager, logger, signService)
	httpServer := sharedApi.NewServer(":80", handler, logger)
	go httpServer.Run()

	<-sharedApp.OnInterrupt()

	fmt.Println("Gracefully shutdown...")
	httpServer.Shutdown()
}

func newJobMap(apiStore api.Store) work.JobMap {
	return work.JobMap{
		jobModel.LedBlink:        job.NewLedBlink(),
		jobModel.MeasureHumidity: job.NewMeasureHumidity(apiStore),
		jobModel.PumpLiquid:      job.NewPumpLiquid(),
	}
}
