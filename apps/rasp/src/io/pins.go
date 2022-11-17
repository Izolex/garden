//go:generate mockgen -source=pins.go -destination=mock/pins.go -package=mock
package io

import (
	"fmt"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"shared/app"
	"shared/app/logger"
	"sync"
)

type Pins interface {
	Get(name string) Pin
}

type pins struct {
	sync.Mutex
	appMode app.Mode
	logger  logger.Logger
	pins    map[string]Pin
}

func NewPins(appMode app.Mode, logger logger.Logger) Pins {
	return &pins{
		appMode: appMode,
		logger:  logger,
		pins:    make(map[string]Pin, 0),
	}
}

func (p *pins) Get(name string) Pin {
	p.Lock()
	defer p.Unlock()

	if pin, ok := p.pins[name]; ok {
		return pin
	}

	pin := p.new(name)
	if pin == nil {
		panic(fmt.Errorf("nothing found on pin %s", name))
	}

	p.pins[name] = pin

	return pin
}

func (p *pins) new(name string) Pin {
	if p.appMode == app.ModeProduction {
		return gpioreg.ByName(name)
	}

	return NewPinFake(name, p.logger)
}
