package periphery

import (
	"fmt"
	"gorm.io/gorm"
	"shared/model/entity/periphery"
	"time"
)

type StateValue struct {
	Pin      string
	Value    float32
	Created  time.Time
	HasValue bool
}
type State map[periphery.Name]StateValue

func (s State) Get(peripheryName periphery.Name) StateValue {
	value, exists := s[peripheryName]
	if !exists {
		panic(fmt.Errorf("periphery %s not exists in state", peripheryName.String()))
	}
	return value
}

func NewState(db *gorm.DB, raspberryId uint) State {
	var values []struct {
		Periphery periphery.Name
		Pin       string
		Value     float32
		CreatedAt time.Time
	}

	result := db.Raw(`
SELECT pinout.periphery_id AS periphery, pin, measurement.value, measurement.created_at AS createdAt
FROM pinout
LEFT JOIN measurement ON measurement.periphery_id = pinout.periphery_id AND pinout.raspberry_id = ?
GROUP BY pinout.periphery_id
ORDER BY measurement.created_at DESC
`, raspberryId).Scan(&values)
	if result.Error != nil {
		panic(result.Error)
	}

	state := State{}
	for _, val := range values {
		state[val.Periphery] = StateValue{val.Pin, val.Value, val.CreatedAt, val.Value > 0}
	}
	return state
}
