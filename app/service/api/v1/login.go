package v1

import (
	userModel "gin-api/app/model/user"
	"gin-api/app/service"
	"gin-api/internal/config"
	"gin-api/internal/constant"
	"gin-api/pkg/jwt"
	"time"
)

const (
	TokenExpireAt = 24*time.Hour
)

// 用户登录
func Login(username, password string) *service.ServiceResponse {
	ok, user := userModel.UserCheckAuth(username, password)
	if ok {
		token, err := jwt.New(string(config.Get().Jwt.Secret)).GenerateSign(int64(user.ID), TokenExpireAt)
		if err != nil {
			return service.Error(constant.StatusJwtGenerateTokenError)
		}

		data := make(map[string]interface{})
		data["id"] = user.ID
		data["token"] = token
		data["expire_at"] = time.Now().Add(TokenExpireAt).Unix()

		return service.Success(data)
	}

	return service.Error(constant.StatusBadRequest)
}
