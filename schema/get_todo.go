package schema

import (
	"github.com/arrizalamin/go-graphql-example/todo"
	"github.com/graphql-go/graphql"
)

var getTodo = &graphql.Field{
	Type: todoType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(string)
		for _, todo := range todos {
			if todo.ID == id {
				return todo, nil
			}
		}
		return todo.Todo{}, nil
	},
}
