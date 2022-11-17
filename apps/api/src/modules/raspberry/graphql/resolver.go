package graphql

import (
	"github.com/graph-gophers/graphql-go"
	peripheryGraphql "main/modules/periphery/graphql"
	"shared/model/entity/raspberry"
	"strconv"
)

type raspberryResolver struct {
	Entity         *raspberry.Entity
	PeripheryList_ []*peripheryGraphql.PeripheryResolver
}

func (r *raspberryResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.Entity.ID)))
}

func (r *raspberryResolver) Name() string {
	return r.Entity.Name
}

func (r *raspberryResolver) IsActive() bool {
	return r.Entity.IsActive
}

func (r *raspberryResolver) PeripheryList() ([]*peripheryGraphql.PeripheryResolver, error) {
	return r.PeripheryList_, nil
}
