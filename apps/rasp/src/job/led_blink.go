package job

import (
	"encoding/json"
	"main/io"
	"main/work"
	"periph.io/x/conn/v3/gpio"
	"shared/model/entity/job"
	"time"
)

func NewLedBlink() work.Func {
	return func(workId work.Id, pins io.Pins, data work.Params, callback work.Callback) work.Stopper {
		var params job.LedBlinkParams
		err := json.Unmarshal(data, &params)
		if err != nil {
			panic("invalid LedBlink job config provided")
		}

		pin := pins.Get(params.Pin)

		err = pin.Out(gpio.High)
		if err != nil {
			panic(err)
		}

		timer := time.AfterFunc(params.Duration, func() {
			err := pin.Out(gpio.Low)
			if err != nil {
				panic(err)
			}
			callback()
		})

		return func() {
			timer.Stop()
			err := pin.Out(gpio.High)
			if err != nil {
				panic(err)
			}
		}
	}
}
