package periphery

type Name int

// keep Ids in sync with DB!
const (
	LiquidPump Name = iota + 1
	HumiditySensor
	Led
)

func (enum Name) String() string {
	return map[Name]string{
		LiquidPump:     "liquidPump",
		HumiditySensor: "humiditySensor",
		Led:            "led",
	}[enum]
}
