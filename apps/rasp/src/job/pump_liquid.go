package job

import (
	"main/io"
	"main/work"
)

func NewPumpLiquid() work.Func {
	return func(workId work.Id, pins io.Pins, params work.Params, callback work.Callback) work.Stopper {
		//c, ok := params.(*job.PumpLiquidParams)
		//if !ok {
		//	panic("invalid water pump job config provided")
		//}
		//
		//pin := pins.Get(c.Pin)
		//
		//err := pin.Out(gpio.Low)
		//if err != nil {
		//	panic(err)
		//}
		//
		//timer := time.AfterFunc(c.Duration, func() {
		//	err := pin.Out(gpio.High)
		//	if err != nil {
		//		panic(err)
		//	}
		//	callback()
		//})

		return func() {
			//timer.Stop()
			//err := pin.Out(gpio.Low)
			//if err != nil {
			//	panic(err)
			//}
		}
	}
}
