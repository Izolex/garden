package schema

import (
	_ "embed"
	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
	raspberryGraphql "main/modules/raspberry/graphql"
)

//go:embed schema.graphql
var gqlSchema string

func NewProvider(db *gorm.DB) *Provider {
	raspberryQueryResolver := raspberryGraphql.NewRaspberryQueryResolver(db)

	resolver := &struct {
		*raspberryGraphql.RaspberryQueryResolver
	}{
		raspberryQueryResolver,
	}

	schema, err := graphql.ParseSchema(gqlSchema, resolver)
	if err != nil {
		panic(err)
	}

	return &Provider{
		schema: schema,
	}
}

type Provider struct {
	schema *graphql.Schema
}

func (p *Provider) GetSchema() *graphql.Schema {
	return p.schema
}
