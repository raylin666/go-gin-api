package logger

import (
	"gin-api/internal/config"
	"gin-api/internal/constant"
	"gin-api/pkg/utils"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"time"
)

const (
	TimestampFormat = "2006-01-02 15:04:05"
)

var (
	// 日志写入实例集合
	loggerWriteMaps map[string]*Logger
)

type Fields logrus.Fields

func (fields Fields) Fields() logrus.Fields {
	return logrus.Fields(fields)
}

type Logger struct {
	// 日志实例
	Instance *logrus.Logger
	// 是否并发写入文件及控制台打印
	IsMulti bool
	// 文件名称(文件写入时存在值)
	FileName string
}

func InitLogger() {
	// 创建文件夹
	utils.CreateDirectory(config.Get().Log.Path)

	// 注册日志写实例
	register()
}

func register() map[string]*Logger {
	loggerWriteMaps = map[string]*Logger{
		constant.LOG_MULTI_APP: instance(constant.LOG_MULTI_APP, true).logger(),
		constant.LOG_APP:       instance(constant.LOG_APP, false).logger(),
		constant.LOG_DB:        instance(constant.LOG_DB, false).logger(),
		constant.LOG_REDIS:     instance(constant.LOG_REDIS, false).logger(),
		constant.LOG_REQUEST:   instance(constant.LOG_REQUEST, false).logger(),
	}

	return loggerWriteMaps
}

// 获取打印日志实例
func New() *logrus.Logger {
	return logrus.StandardLogger()
}

// 获取日志写入文件实例
func NewWrite(filename string) *logrus.Logger {
	var (
		logger *Logger
		ok     bool
	)

	if logger, ok = loggerWriteMaps[filename]; !ok {
		return New()
	}

	return logger.Instance
}

// 创建新的实例
func instance(filename string, multi bool) *Logger {
	return &Logger{
		Instance: logrus.New(),
		FileName: filename,
		IsMulti:  multi,
	}
}

// Logger 实例初始化配置
func (logger *Logger) logger() *Logger {
	// 设置日志级别
	logger.Instance.SetLevel(logrus.DebugLevel)
	// 设置日志格式
	logger.Instance.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TimestampFormat,
	})

	file := path.Join(config.Get().Log.Path, logger.FileName)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		file+"-%Y-%m-%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(file),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err == nil {
		if logger.IsMulti {
			logger.Instance.SetOutput(io.MultiWriter(os.Stdout, logWriter))
		} else {
			logger.Instance.SetOutput(logWriter)
		}
	} else {
		log.Printf("(%s) failed to create rotatelogs: %s", logger.FileName, err)
	}

	return logger
}
