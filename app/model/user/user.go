package user

import (
	"gin-api/app/model"
	"gin-api/pkg/database"
)

type User struct {
	model.Model

	Username string `json:"username"`
	Avatar	 string `json:"avatar"`
	Phone    string `json:"phone"`
}

// 获取单个用户信息
func GetUserOne(id int) (user []User) {
	database.GetDB("local").Where("id = ?", id).First(&user)
	return
}