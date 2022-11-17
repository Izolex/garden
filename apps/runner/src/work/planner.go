package work

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"shared/app/logger"
	"shared/model/entity/plan"
	"shared/model/entity/raspberry"
	"time"
)

type Plan func(raspberry.Entity)
type PlanMap map[plan.Name]Plan

type Planner interface {
	Run()
}

type planner struct {
	tickerDuration time.Duration
	db             *gorm.DB
	logger         logger.Logger
	plans          PlanMap
}

func NewPlanner(tickerDuration int, db *gorm.DB, logger logger.Logger, plans PlanMap) Planner {
	return &planner{time.Duration(tickerDuration) * time.Millisecond, db, logger, plans}
}

func (p *planner) Run() {
	ticker := time.NewTicker(p.tickerDuration)
	defer ticker.Stop()

	for {
		<-ticker.C

		var raspberries []raspberry.Entity

		result := p.db.Find(&raspberries)
		if result.Error != nil {
			panic(result.Error)
		}

		for _, rasp := range raspberries {
			raspPlan, exists := p.plans[plan.Name(rasp.PlanId)]
			if !exists {
				panic(fmt.Errorf("plan id %d does not exists", rasp.PlanId))
			}
			p.runPlan(raspPlan, rasp)
		}
	}
}

func (p *planner) runPlan(plan Plan, rasp raspberry.Entity) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			p.logger.Error(errors.New("plan failed"))
		}
	}()
	plan(rasp)
}
