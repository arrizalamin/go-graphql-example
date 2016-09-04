package schema

import "github.com/arrizalamin/go-graphql-example/todo"

func convertPointersToTypes(array []*todo.Todo) (todos []todo.Todo) {
	for _, t := range array {
		todos = append(todos, *t)
	}
	return
}
