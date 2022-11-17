package work

import (
	"main/periphery"
	"main/raspberry"
	jobModel "shared/model/entity/job"
	peripheryModel "shared/model/entity/periphery"
	raspberryModel "shared/model/entity/raspberry"
	"time"
)

func MakeVisualInspection(raspService raspberry.Service, peripheryState periphery.State, rasp raspberryModel.Entity) {
	led := peripheryState.Get(peripheryModel.Led)
	raspService.Work(rasp, jobModel.LedBlink, jobModel.NewLedBlinkParams(led.Pin, 3*time.Second))
}
