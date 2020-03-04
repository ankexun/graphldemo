package schema

import (
	"errors"
	"graphqldemo/models"
	"graphqldemo/utils"
	"log"

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
// 用户信息
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

var loginType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Login",
	Description: "Login",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// 登录,返回token
var login = graphql.Field{
	Name:        "login",
	Description: "login",
	Type:        loginType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"test": &graphql.ArgumentConfig{
			Type:         graphql.String,
			DefaultValue: "test",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username, ok1 := p.Args["username"].(string)
		password, ok2 := p.Args["password"].(string)
		test, ok3 := p.Args["test"].(string)
		if ok3 {
			log.Println(test)
		}
		if !ok1 || !ok2 {
			return nil, errors.New("missing required arguments: username or password. ")
		}

		user := models.User{Username: username, Password: password}
		if _, err := user.Authenticate(); err != nil {
			return nil, err
		}

		result := make(map[string]string)
		var err error
		result["token"], err = utils.GenerateToken(username, password)

		return result, err
	},
}
