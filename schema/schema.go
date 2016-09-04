package schema

import (
	"log"

	"github.com/graphql-go/graphql"
)

var Schema = createSchema()

func createSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}
