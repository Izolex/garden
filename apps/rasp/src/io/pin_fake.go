package io

import (
	"fmt"
	"periph.io/x/conn/v3/gpio"
	"shared/app/logger"
)

type pinFake struct {
	name   string
	logger logger.Logger
}

func NewPinFake(name string, logger logger.Logger) Pin {
	return &pinFake{
		name:   name,
		logger: logger,
	}
}

func (p *pinFake) Out(l gpio.Level) error {
	if p.logger != nil {
		p.logger.Info(fmt.Sprintf("pin %s %t\n", p.name, l))
	}
	return nil
}
