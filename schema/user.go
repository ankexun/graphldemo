package schema

import (
	"errors"
	"graphqldemo/models"

	"github.com/graphql-go/graphql"
)

// 定义查询对象的字段，支持嵌套
var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "User Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// 定义query
// 用户登录
var queryUser = graphql.Field{
	Name:        "User Info",
	Description: "User Info",
	Type:        userType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		id, ok := p.Args["id"].(int)
		if !ok {
			return nil, errors.New("missing required arguments: id. ")
		}

		result, err = new(models.User).GetUserByID(id)
		if err != nil {
			return nil, err
		}

		return result, err
	},
}
