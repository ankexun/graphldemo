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

// 定义mutation,增删改操作
// add
var addArticle = graphql.Field{
	Name:        "新文章",
	Description: "增加新文章",
	Type:        articleType,
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		title, _ := p.Args["title"].(string)
		content, _ := p.Args["content"].(string)

		result, err = models.CreateNewArticle(title, content)
		if err != nil {
			return nil, err
		}

		return result, nil
	},
}

// update
var updateArticle = graphql.Field{
	Name:        "编辑文章",
	Description: "编辑文章",
	Type:        articleType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		id, _ := p.Args["id"].(int)
		title, _ := p.Args["title"].(string)
		content, _ := p.Args["content"].(string)

		result, err = models.UpdateArticle(id, title, content)
		if err != nil {
			return nil, err
		}

		return result, nil
	},
}

// delete
var deleteArticle = graphql.Field{
	Name:        "删除文章",
	Description: "删除指定Id的文章",
	Type:        articleType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		id, _ := p.Args["id"].(int)

		// 查找文章是否存在
		result, err = models.GetArticleByID(id)
		if err != nil {
			return nil, err
		}

		if err = models.DeleteArticle(id); err != nil {
			return nil, err
		}

		return result, nil
	},
}

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
