package work

import (
	"shared/model/database"
	"time"
)

type Entity struct {
	ID          uint
	RaspberryId uint
	JobId       uint
	Params      database.JSON
	CreatedAt   time.Time
}

func (Entity) TableName() string {
	return "work"
}
