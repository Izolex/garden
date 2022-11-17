package raspberry

type Entity struct {
	ID       uint
	Name     string
	Address  string
	IsActive bool
	PlanId   uint
	User     string
	Password string
}

func (Entity) TableName() string {
	return "raspberry"
}
