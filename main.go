package main

import (
	"context"
	"graphqldemo/schema"
	"graphqldemo/utils"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/graphql-go/handler"
)

// main
func main() {
	h := Register()

	// 跨域
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		// Debug:            true,
	})

	http.Handle("/graphql", c.Handler(h))
	http.HandleFunc("/login", utils.CreateTokenEndpoint)
	log.Println("Now server is running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

// 初始化handler
func Register() http.Handler {
	h := handler.New(&handler.Config{
		Schema:     &schema.Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	// token传参
	hdl := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "token", r.Header.Get("token"))
		h.ContextHandler(ctx, w, r)
	})

	return hdl
}
