//go:generate mockgen -source=pin.go -destination=mock/pin.go -package=mock
package io

import (
	"periph.io/x/conn/v3/gpio"
)

// gpio.PinIO
type Pin interface {
	Out(l gpio.Level) error
}
