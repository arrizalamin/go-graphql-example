package schema

import (
	"github.com/arrizalamin/go-graphql-example/todo"
	"github.com/graphql-go/graphql"
)

var completeTodo = &graphql.Field{
	Type: todoType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(string)

		var task *todo.Todo
		for _, t := range todos {
			if t.ID == id {
				t.Done = true
				task = t
			}
		}

		return *task, nil
	},
}
