package schema

import (
	"github.com/arrizalamin/go-graphql-example/todo"
	"github.com/graphql-go/graphql"
)

var deleteTodo = &graphql.Field{
	Type: todoType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(string)

		var newTodos []*todo.Todo
		var task *todo.Todo
		for _, t := range todos {
			if t.ID == id {
				task = t
			} else {
				newTodos = append(newTodos, t)
			}
		}
		todos = newTodos

		return *task, nil
	},
}
