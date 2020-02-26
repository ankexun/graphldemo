package schema

import (
	"errors"
	"graphqldemo/models"

	"github.com/graphql-go/graphql"
)

// 定义查询对象的字段，支持嵌套
var articleType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Article",
	Description: "Article Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// 处理查询请求,获取文章列表
var queryArticles = graphql.Field{
	Name:        "QueryArticles",
	Description: "Query Article List",
	Type:        graphql.NewList(articleType),
	// 无需处理参数
	// Resolve是一个处理请求的函数，具体处理逻辑可在此进行
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		// 调用Article的models里面的方法查询数据
		result = models.GetAllArticles()

		return result, nil
	},
}

// 查询单篇文章
var queryArticle = graphql.Field{
	Name:        "QueryArticle",
	Description: "Query Article",
	Type:        articleType,
	// Args是定义在GraphQL查询中支持的查询字段，
	// 可自行随意定义，如加上limit,start这类
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	// Resolve是一个处理请求的函数，具体处理逻辑可在此进行
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		// Args里面定义的字段在p.Args里面，对应的取出来
		// 因为是interface{}的值，需要类型转换
		id, ok := p.Args["id"].(int)
		if !ok {
			return nil, errors.New("missing required arguments: id. ")
		}

		result, err = models.GetArticleByID(id)

		// 调用Hello这个model里面的Query方法查询数据
		return result, err
	},
}

// 定义根查询节点及各种查询
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"articles": &queryArticles, //分别对应前面定义的query
		"article":  &queryArticle,
	},
})

// 定义Schema用于http handler处理
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: nil,
})
