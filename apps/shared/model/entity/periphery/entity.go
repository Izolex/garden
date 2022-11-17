package periphery

type Entity struct {
	ID           uint
	Name         string
	IsMeasurable bool
}

func (Entity) TableName() string {
	return "periphery"
}
