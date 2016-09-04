package schema

import (
	"github.com/arrizalamin/go-graphql-example/todo"

	"github.com/graphql-go/graphql"
)

var todos = []*todo.Todo{}

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"new":      newTodo,
		"complete": completeTodo,
		"delete":   deleteTodo,
	},
})

func init() {
	for _, name := range []string{"a", "b", "c"} {
		task := todo.New(name)
		todos = append(todos, &task)
	}
}
