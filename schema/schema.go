package schema

import "github.com/graphql-go/graphql"

// query
// 定义根查询节点及各种查询
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"articles": &queryArticles,
		"article":  &queryArticle,
	},
})

// mutation
// 定义增删改方法
var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "mutation",
	Description: "增删改",
	Fields: graphql.Fields{
		"add":    &addArticle,
		"update": &updateArticle,
		"delete": &deleteArticle,
	},
})

// 定义Schema用于http handler处理
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: mutationType,
})
