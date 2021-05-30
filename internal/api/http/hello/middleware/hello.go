package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Hello 服务打印中间件
func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hello GO. I'm Middleware.")
	}
}
