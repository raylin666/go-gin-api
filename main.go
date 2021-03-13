package main

import (
	"context"
	"fmt"
	"gin-api/internal/config"
	"gin-api/internal/constant"
	"gin-api/internal/env"
	"gin-api/internal/routers"
	"gin-api/pkg/cache"
	"gin-api/pkg/database"
	"gin-api/pkg/logger"
	"gin-api/pkg/shutdown"
	"net/http"
	"time"
)

func init() {
	// 配置初始化
	config.InitConfig()
	// 环境初始化
	env.InitEnv()
	// 日志初始化
	logger.InitLogger()
	// 数据库初始化
	database.InitDatabase()
	// 缓存初始化
	cache.InitRedis()
}

func main() {
	router := routers.InitRouter()

	host := config.Get().Http.Host
	port := config.Get().Http.Port

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", host, port),
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.NewWrite(constant.LOG_APP).Info(fmt.Sprintf("start http server listening %s:%d", host, port))

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.NewWrite(constant.LOG_MULTI_APP).Fatal("http server startup err", err)
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 HTTP 服务
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				logger.NewWrite(constant.LOG_MULTI_APP).Error("server shutdown err", err)
			} else {
				logger.NewWrite(constant.LOG_MULTI_APP).Info("server shutdown success")
			}
		},

		// 关闭 Database
		func() {
			if err := database.CloseAll(); err != nil {
				logger.NewWrite(constant.LOG_MULTI_APP).Error("database close err", err)
			} else {
				logger.NewWrite(constant.LOG_MULTI_APP).Info("database close success")
			}
		},

		// 关闭缓存
		func() {
			if err := cache.CloseAll(); err != nil {
				logger.NewWrite(constant.LOG_MULTI_APP).Error("cache close err", err)
			} else {
				logger.NewWrite(constant.LOG_MULTI_APP).Info("cache close success")
			}
		},
	)
}
