package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"time"
)

type PeripheryValueResolver struct {
	Value_    float64
	DateTime_ time.Time
}

func (r *PeripheryValueResolver) Value() float64 {
	return r.Value_
}

func (r *PeripheryValueResolver) DateTime() graphql.Time {
	return graphql.Time{Time: r.DateTime_}
}
