package schema

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// query
// 定义根查询节点及各种查询
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"queryArticles":  &queryArticles,
		"getArticle":     &queryArticle,
		"login":          &login,
		"queryTodolists": &queryTodoLists,
	},
})

// mutation
// 定义增删改方法
var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "mutation",
	Description: "增删改",
	Fields: graphql.Fields{
		"addArticle":     &addArticle,
		"editArticle":    &updateArticle,
		"deleteArticle":  &deleteArticle,
		"addTodolist":    &addTodoList,
		"updateTodolist": &updateTodoList,
	},
})

// 定义Schema用于http handler处理
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: mutationType,
})

// 初始化handler
func Register() http.Handler {
	h := handler.New(&handler.Config{
		Schema:     &schema,
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
