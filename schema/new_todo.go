package schema

import (
	"github.com/arrizalamin/go-graphql-example/todo"
	"github.com/graphql-go/graphql"
)

var newTodo = &graphql.Field{
	Type: todoType,
	Args: graphql.FieldConfigArgument{
		"text": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		text, _ := p.Args["text"].(string)
		task := todo.New(text)
		todos = append(todos, &task)
		return task, nil
	},
}
