package initx

import (
	"github.com/raylin666/go-gin-api/config"
	"github.com/raylin666/go-gin-api/environment"
	"github.com/raylin666/go-gin-api/pkg/cache"
	"github.com/raylin666/go-gin-api/pkg/database"
	"github.com/raylin666/go-gin-api/pkg/logger"
)

// 初始化应用配置
type Config struct {
	YmlEnvFileName string
}

// 初始化应用
func InitApplication(c *Config) {
	// 初始化加载配置文件
	config.InitAutoloadConfig(c.YmlEnvFileName)
	// 初始化环境
	environment.InitEnvironment()
	// 日志初始化
	logger.InitLogger()
	// 数据库初始化
	database.InitDB()
	// 缓存初始化
	cache.InitRedis()
}
