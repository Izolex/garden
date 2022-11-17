package measurement

import "time"

type Entity struct {
	ID          uint
	WorkId      uint
	PeripheryId uint
	Value       float32
	CreatedAt   time.Time
}

func (Entity) TableName() string {
	return "measurement"
}
