package initx

import (
	"go-gin-api/config"
	"go-gin-api/internal/constant"
	"go-gin-api/internal/environment"
	"go-gin-api/pkg/cache"
	"go-gin-api/pkg/database"
	"go-gin-api/pkg/logger"
)

type InitApp struct {
	EnvFileName string
}

func NewInitApp(envFileName string) *InitApp {
	return &InitApp{
		EnvFileName: envFileName,
	}
}

func (app *InitApp) Run()  {
	if app.EnvFileName == "" {
		app.EnvFileName = constant.EnvFileName
	}

	// 初始化加载配置文件
	config.InitAutoloadConfig(app.EnvFileName)
	// 初始化环境
	environment.InitEnvironment()
	// 日志初始化
	logger.InitLogger()
	// 数据库初始化
	database.InitDatabase()
	// 缓存初始化
	cache.InitRedis()
}
