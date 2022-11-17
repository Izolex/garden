package job

import (
	"shared/model/entity/periphery"
	"time"
)

func NewLedBlinkParams(pin string, duration time.Duration) LedBlinkParams {
	return LedBlinkParams{pin, duration}
}

type LedBlinkParams struct {
	Pin      string        `json:"pin"`
	Duration time.Duration `json:"duration"`
}

func NewMeasureHumidityParams(pin string, periphery periphery.Name) MeasureHumidityParams {
	return MeasureHumidityParams{pin, periphery}
}

type MeasureHumidityParams struct {
	Pin       string         `json:"pin"`
	Periphery periphery.Name `json:"periphery"`
}

func NewPumpLiquidParams(pin string, duration time.Duration) PumpLiquidParams {
	return PumpLiquidParams{pin, duration}
}

type PumpLiquidParams struct {
	Pin      string        `json:"pin"`
	Duration time.Duration `json:"duration"`
}
