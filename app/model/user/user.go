package user

import (
	"gin-api/app/model"
	"gin-api/pkg/utils"
)

type User struct {
	model.Model

	Username string `json:"username"`
	Password string `json:"-"`
	Avatar	 string `json:"avatar"`
	Phone    string `json:"phone"`
}

// 验证用户并获取当前用户信息
func UserCheckAuth(username, password string) (ok bool, user User) {
	model.GetLocalDB().Where(User{Username: username}).First(&user)
	if user.ID > 0 && utils.PasswordVerify(password, user.Password) {
		return true, user
	}

	return false, user
}

// 获取单个用户信息
func GetUserOne(id uint64) (user []User) {
	model.GetLocalDB().Where("id = ?", id).First(&user)
	return
}

