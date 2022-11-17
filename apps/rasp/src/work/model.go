package work

import (
	"main/io"
	"shared/model/entity/job"
)

type Id uint
type Func func(Id, io.Pins, Params, Callback) Stopper
type Params []byte
type Callback func()
type Stopper func()
type JobMap map[job.Name]Func
