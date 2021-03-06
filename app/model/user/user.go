package user

import (
	"gin-api/app/model"
	"gin-api/pkg/database"
	"gin-api/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	model.Model

	Username string `json:"username"`
	Password string `json:"-"`
	Avatar	 string `json:"avatar"`
	Phone    string `json:"phone"`
}

func getDB() *gorm.DB {
	return database.GetDB("local")
}

// 验证用户并获取当前用户信息
func UserCheckAuth(username, password string) (ok bool, user User) {
	getDB().Where(User{Username: username}).First(&user)
	if user.ID > 0 && utils.PasswordVerify(password, user.Password) {
		return true, user
	}

	return false, user
}

// 获取单个用户信息
func GetUserOne(id uint64) (user []User) {
	getDB().Where("id = ?", id).First(&user)
	return
}

