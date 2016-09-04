package schema

import (
	"github.com/arrizalamin/go-graphql-example/todo"

	"github.com/graphql-go/graphql"
)

var todos = []*todo.Todo{}

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"new": &graphql.Field{
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
		},
		"complete": &graphql.Field{
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
		},
		"delete": &graphql.Field{
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
		},
	},
})

func init() {
	for _, name := range []string{"a", "b", "c"} {
		task := todo.New(name)
		todos = append(todos, &task)
	}
}
