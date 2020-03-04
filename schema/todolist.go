package schema

import (
	"errors"
	"graphqldemo/models"
	"graphqldemo/utils"
	"log"

	"github.com/graphql-go/graphql"
)

// todolist
var todolistType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "todolist",
	Description: "todolist",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"user_id": &graphql.Field{
			Type: graphql.Int,
		},
		"status": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// 查询列表
var queryTodoLists = graphql.Field{
	Name:        "QueryTodolists",
	Description: "Query Todo List",
	Type:        graphql.NewList(todolistType),
	Args: graphql.FieldConfigArgument{
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page_num": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page_size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		log.Println(p.Info.FieldName)
		if err := utils.ValidateJWT(p.Context.Value("token").(string)); err != nil {
			log.Println(err)
			return nil, err
		}
		userId, ok1 := p.Args["user_id"].(int)
		pageNum, ok2 := p.Args["page_num"].(int)
		pageSize, ok3 := p.Args["page_size"].(int)
		if !ok1 || !ok2 || !ok3 {
			return nil, errors.New("missing required arguments: user_id or page_num or page_size. ")
		}

		result, err = new(models.ToDoList).GetTodoLists(int64(userId), pageNum, pageSize)
		return
	},
}

// 定义mutation,增删改操作
// add
var addTodoList = graphql.Field{
	Name:        "新增todolist",
	Description: "新增todolist",
	Type:        todolistType,
	Args: graphql.FieldConfigArgument{
		"content": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"user_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		if err := utils.ValidateJWT(p.Context.Value("token").(string)); err != nil {
			log.Println(err)
			return nil, err
		}

		content, ok1 := p.Args["content"].(string)
		userId, ok2 := p.Args["user_id"].(int)
		if !ok1 || !ok2 {
			return nil, errors.New("missing required arguments: content or user_id. ")
		}
		todo := models.ToDoList{
			Content: content,
			UserId:  int64(userId),
		}
		_, err := todo.Insert()

		return nil, err
	},
}

// update
var updateTodoList = graphql.Field{
	Name:        "修改todolist",
	Description: "修改todolist状态",
	Type:        todolistType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		if err := utils.ValidateJWT(p.Context.Value("token").(string)); err != nil {
			log.Println(err)
			return nil, err
		}

		id, ok1 := p.Args["id"].(int)
		status, ok2 := p.Args["status"].(int)
		if !ok1 || !ok2 {
			return nil, errors.New("missing required arguments: id or status. ")
		}

		todo := models.ToDoList{
			Id:     int64(id),
			Status: status,
		}
		_, err := todo.Update()

		return nil, err
	},
}
