package job

type Name int

// keep Ids in sync with DB!
const (
	LedBlink Name = iota + 1
	MeasureHumidity
	PumpLiquid
)
