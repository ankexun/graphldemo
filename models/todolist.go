package models

import (
	"bytes"
	"database/sql"
	"log"
	"time"
)

// todolist
// 同时使用xorm的tag和json的tag,json必须放在前面
type ToDoList struct {
	Id        int64  //主键ID
	Content   string //内容
	UserId    int64  `json:"user_id" xorm: "user_id"`       //对应用户ID
	CreatedAt Time   `json:"created_at" xorm: "created_at"` //创建时间
	UpdatedAt Time   `json:"updated_at" xorm: "updated_at"` //最后更新时间
	Status    int    //0：未执行，1：已完成，2：删除
}

/****
* 根据用户ID、分页参数获取纪录
* PageNum从0开始
 */
func (todo *ToDoList) GetTodoLists(userId int64, pageNum, pageSize int) ([]*ToDoList, error) {

	/* `SELECT * FROM (
	SELECT * FROM (SELECT * FROM todo_list WHERE user_id = ? AND status = 0 ORDER BY id DESC LIMIT 999999999) AS a
	UNION ALL
	SELECT * FROM (SELECT * FROM todo_list WHERE user_id = ? AND (status = 1 OR status = 2) ORDER BY updated_at DESC LIMIT 999999999) AS b
	) AS c LIMIT ?,?`
	*/
	var sqlBuffer bytes.Buffer
	var args []interface{}
	sqlBuffer.WriteString("SELECT * FROM ( ")
	sqlBuffer.WriteString("SELECT * FROM (SELECT * FROM todo_list WHERE user_id = ? AND status = 0 ORDER BY id DESC LIMIT 999999999) AS a  ")
	args = append(args, userId)
	sqlBuffer.WriteString("UNION ALL ")
	sqlBuffer.WriteString("SELECT * FROM (SELECT * FROM todo_list WHERE user_id = ? AND (status = 1 OR status = 2) ORDER BY updated_at DESC LIMIT 999999999) as b ")
	args = append(args, userId)
	sqlBuffer.WriteString(") AS c LIMIT ?,?")
	args = append(args, pageNum, pageSize)

	var result []*ToDoList
	if err := DbRead.SQL(sqlBuffer.String(), args...).Find(&result); err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

/**
* 添加
 */
func (todo *ToDoList) Insert() (sql.Result, error) {

	result, err := DbWrite.Exec("INSERT INTO todo_list(content,user_id,status,created_at,updated_at) VALUES (?,?,?,?,?)",
		todo.Content, todo.UserId, 0, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

/***
Update 修改状态
*/
func (todo *ToDoList) Update() (sql.Result, error) {
	result, err := DbWrite.Exec("UPDATE todo_list SET status=?, updated_at=? WHERE id=?",
		todo.Status, time.Now(), todo.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

/*
获取记录总数
*/
func (todo *ToDoList) Total() (int, error) {

	count, err := DbRead.Count(todo)

	if err != nil {
		log.Println(err)
		return -1, err
	}

	return int(count), nil
}
