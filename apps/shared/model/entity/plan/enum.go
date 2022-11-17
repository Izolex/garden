package plan

type Name int

// keep Ids in sync with DB!
const (
	MungBeanSprouts Name = iota + 1
)

func (enum Name) String() string {
	return map[Name]string{
		MungBeanSprouts: "mungBeanSprouts",
	}[enum]
}
