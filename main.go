package main

import (
	"fmt"
	"graphqldemo/schema"
	"net/http"

	"github.com/rs/cors"

	"github.com/graphql-go/handler"
)

// main
func main() {
	h := Register()
	// 解决跨域的问题
	hdl := cors.Default().Handler(h)
	http.Handle("/graphql", hdl)
	fmt.Println("Now server is running on port 9090")
	http.ListenAndServe(":9090", nil)
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
