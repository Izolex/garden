package job

import (
	"encoding/json"
	"fmt"
	"main/api"
	"main/io"
	"main/work"
	"math/rand"
	"shared/model/entity/job"
	"strconv"
)

func NewMeasureHumidity(store api.Store) work.Func {
	return func(workId work.Id, pins io.Pins, data work.Params, callback work.Callback) work.Stopper {
		var params job.MeasureHumidityParams
		err := json.Unmarshal(data, &params)
		if err != nil {
			panic("invalid humidity job config provided")
		}

		randStr := fmt.Sprintf("%.2f", rand.Float32())
		val, err := strconv.ParseFloat(randStr, 32)
		if err != nil {
			panic(err)
		}

		store.Measurement(workId, params.Periphery, float32(val))

		return func() {}
	}
}
