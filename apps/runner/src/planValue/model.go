package planValue

import "gorm.io/gorm"

type Model interface {
	GetValue(raspberryId uint, valueName string) (float32, error)
}

func NewModel(db *gorm.DB) Model {
	return &model{db: db}
}

type row struct {
	Value float32
}

type model struct {
	db *gorm.DB
}

func (m *model) GetValue(raspberryId uint, valueName string) (float32, error) {
	var valueRow row

	result := m.db.Raw(`
SELECT value
FROM raspberry_plan_value
INNER JOIN plan_value ON plan_value.plan_id = raspberry_plan_value.plan_value_id AND plan_value.name = ?
WHERE raspberry_plan_value.raspberry_id = ?
`, valueName, raspberryId).Scan(&valueRow)
	if result.Error != nil {
		return 0, result.Error
	}

	return valueRow.Value, nil
}
