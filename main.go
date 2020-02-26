package main

import (
	"fmt"
	"graphqldemo/schema"
	"net/http"

	"github.com/graphql-go/handler"
)

// main
func main() {
	h := Register()
	http.Handle("/graphql", h)
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

// 初始化handler
func Register() *handler.Handler {
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return h
}
