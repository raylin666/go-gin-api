package main

import (
	"go-gin-api/config"
	"go-gin-api/internal/api/router"
	"go-gin-api/internal/api/server"
	"go-gin-api/internal/constant"
	"go-gin-api/internal/environment"
	"go-gin-api/pkg/cache"
	"go-gin-api/pkg/database"
	"go-gin-api/pkg/logger"
)

func init()  {
	// 初始化加载配置文件
	config.InitAutoloadConfig(constant.EnvFileName)
	// 初始化环境
	environment.InitEnvironment()
	// 日志初始化
	logger.InitLogger()
	// 数据库初始化
	database.InitDatabase()
	// 缓存初始化
	cache.InitRedis()
}

func main()  {
	r := &router.Router{}
	server.NewHttpServer(r)
}
