package main

import (
	"net/http"

	"github.com/arrizalamin/go-graphql-example/schema"
	"github.com/fvbock/endless"
	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)
	endless.ListenAndServe("localhost:8080", nil)
}
