package plan

import (
	"gorm.io/gorm"
	"main/periphery"
	"main/planValue"
	"main/raspberry"
	"main/work"
	jobModel "shared/model/entity/job"
	peripheryModel "shared/model/entity/periphery"
	raspberryModel "shared/model/entity/raspberry"
	"time"
)

func NewMungBeanSprouts(db *gorm.DB, raspService raspberry.Service, planValueModel planValue.Model) work.Plan {
	return func(rasp raspberryModel.Entity) {
		peripheryState := periphery.NewState(db, rasp.ID)
		workModel := work.NewModel(db)

		work.MakeVisualInspection(raspService, peripheryState, rasp)

		liquidPump := peripheryState.Get(peripheryModel.LiquidPump)
		lastPumpLiquid, err := workModel.GetLast(rasp.ID, jobModel.PumpLiquid)
		if err != nil {
			panic(err)
		}

		liquidPumpInterval, err := planValueModel.GetValue(rasp.ID, "liquidPumpInterval")
		if err != nil {
			panic(err)
		}
		liquidPumpDuration, err := planValueModel.GetValue(rasp.ID, "liquidPumpDuration")
		if err != nil {
			panic(err)
		}

		before := time.Now().Truncate(time.Duration(liquidPumpInterval) * time.Hour)
		if lastPumpLiquid == nil || lastPumpLiquid.CreatedAt.Before(before) {
			params := jobModel.NewLedBlinkParams(liquidPump.Pin, time.Duration(liquidPumpDuration)*time.Second)
			raspService.Work(rasp, jobModel.PumpLiquid, params)
		}
	}
}
