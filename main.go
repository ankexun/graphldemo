package main

import (
	"context"
	"graphqldemo/schema"
	"graphqldemo/utils"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

// main
func main() {
	h := Register()

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "token", r.Header.Get("token"))

		// 解决跨域的问题
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		// w.Header().Set("content-type", "application/json")             //返回数据格式是json

		h.ContextHandler(ctx, w, r)
	})
	http.HandleFunc("/login", utils.CreateTokenEndpoint)
	log.Println("Now server is running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

// 初始化handler
func Register() *handler.Handler {
	h := handler.New(&handler.Config{
		Schema:     &schema.Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	return h
}
