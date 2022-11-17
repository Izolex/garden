package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"shared/model/entity/periphery"
	"strconv"
)

type PeripheryResolver struct {
	Entity *periphery.Entity
	Values []*PeripheryValueResolver
}

func (r *PeripheryResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.Entity.ID)))
}

func (r *PeripheryResolver) Name() string {
	return r.Entity.Name
}

func (r *PeripheryResolver) IsMeasurable() bool {
	return r.Entity.IsMeasurable
}

func (r *PeripheryResolver) ValueList() ([]*PeripheryValueResolver, error) {
	return r.Values, nil
}
