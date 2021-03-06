package jwt

import (
	userModel "gin-api/app/model/user"
	"gin-api/internal/config"
	"gin-api/internal/constant"
	"gin-api/internal/http"
	"gin-api/internal/utils"
	"gin-api/pkg/jwt"
	"time"
)

func JWT() utils.ContextHandlerFunc {
	return func(ctx *utils.Context) {
		var code = constant.StatusOK

		token := ctx.Request.Header.Get("token")
		if len(token) == 0 {
			code = constant.StatusBadRequest
		} else {
			claims, err := jwt.New(string(config.Get().Jwt.Secret)).ParseSign(token)
			if err != nil {
				code = constant.StatusJwtAuthCheckTokenFailError
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = constant.StatusJwtCheckTokenTimeoutError
			}

			if claims != nil {
				user := userModel.GetUserOne(uint64(claims.UserID))
				if user == nil {
					code = constant.StatusUserNotFoundError
				} else {
					utils.SetContextUser(ctx, user)
				}
			}
		}

		if code != constant.StatusOK {
			http.ErrorResponse(ctx, code)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}