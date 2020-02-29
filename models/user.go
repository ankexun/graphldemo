package models

import "errors"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var accountsMock = []User{
	User{
		Id:       1,
		Username: "admin",
		Password: "1234",
	},
	User{
		Id:       2,
		Username: "guest",
		Password: "5678",
	},
}

// Get User Info
func (u *User) GetUserByID(id int) (*User, error) {
	for _, user := range accountsMock {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

// 校验用户名与密码
func (u *User) Authenticate() (*User, error) {
	for _, user := range accountsMock {
		if user.Username == u.Username && user.Password == u.Password {
			return &user, nil
		}
	}

	return nil, errors.New("User not found or password is wrong")
}
