package logger

import (
	"github.com/raylin666/go-utils/logger"
	"go-gin-api/config"
	"go-gin-api/internal/constant"
)

var (
	loggerConfig = new(logger.LoggerConfig)
)

// 日志初始化 (注册日志文件)
func InitLogger() {
	loggerConfig.LogPath = config.Get().Logs.Path
	loggerConfig.RegisterWriteFileNameArr = []string{
		constant.LogApp,
		constant.LogSql,
		constant.LogRequest,
		constant.LogDb,
		constant.LogCache,
	}
	loggerConfig.InitLogger()
}
