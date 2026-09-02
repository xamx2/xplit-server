package gql

import (
	_ "embed"

	"github.com/graph-gophers/graphql-go"
	"github.com/xamx2/xplit-server/gql/resolver"
)

//go:embed schema.gql
var gqlSchema string

func NewSchema() *graphql.Schema {
	return graphql.MustParseSchema(gqlSchema, &resolver.RootResolver{},
		graphql.DisableFieldSelections(),
		graphql.UseFieldResolvers(),
	)
}
