package schema

import "github.com/graphql-go/graphql"

var allTodos = &graphql.Field{
	Type: graphql.NewList(todoType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return convertPointersToTypes(todos), nil
	},
}
