package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/raylin666/go-gin-api/config"
	"github.com/raylin666/go-gin-api/constant"
	"github.com/raylin666/go-gin-api/pkg/utils"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var (
	WriteMaps map[string]*Logger
)

type H logrus.Fields

func (data H) Fields() logrus.Fields {
	return logrus.Fields(data)
}

type Logger struct {
	// 日志实例
	Instance *logrus.Logger
	// 文件名称(文件写入时存在值)
	FileName string
	// 日志级别
	Level logrus.Level
	// 日志格式
	Format logrus.Formatter
	// 是否并发写入文件及控制台打印
	Multi bool
}

func InitLogger() {
	var logger *Logger

	// 创建文件夹
	utils.CreateDirectory(config.Get().Logs.Path)

	// 注册日志写实例
	RegisterWriteMaps(map[string]*Logger{
		constant.LOG_MULTI_APP: logger.instanceMulti(constant.LOG_MULTI_APP),
		constant.LOG_APP:       logger.instance(constant.LOG_APP),
		constant.LOG_DB:        logger.instance(constant.LOG_DB),
		constant.LOG_REDIS:     logger.instance(constant.LOG_REDIS),
		constant.LOG_REQUEST:   logger.instance(constant.LOG_REQUEST),
		constant.LOG_SQL:   	logger.instance(constant.LOG_SQL),
	})
}

// 注意: 该注册方法必须在服务启动前调用, 否则会有问题
func RegisterWriteMaps(maps map[string]*Logger) map[string]*Logger {
	for filename, logger := range maps {
		WriteMaps[filename] = logger
	}
	return WriteMaps
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

	if logger, ok = WriteMaps[filename]; !ok {
		return New()
	}

	return logger.Instance
}

func (logger *Logger) instance(filename string) *Logger {
	logger.FileName = filename
	return logger.create()
}

func (logger *Logger) instanceMulti(filename string) *Logger {
	logger.FileName = filename
	logger.Multi = true
	return logger.create()
}

// 创建 Logger 实例 初始化配置
func (logger *Logger) create() *Logger {
	defer logger.reset()

	l := logrus.New()

	// 设置日志级别
	if logger.Level == 0 {
		l.SetLevel(logrus.DebugLevel)
	} else {
		l.SetLevel(logger.Level)
	}

	// 设置日志格式
	if logger.Format == nil {
		l.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: constant.TIMESTAMP_FORMAT,
		})
	} else {
		l.SetFormatter(logger.Format)
	}

	file := path.Join(config.Get().Logs.Path, logger.FileName)
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
		if logger.Multi {
			l.SetOutput(io.MultiWriter(os.Stdout, logWriter))
		} else {
			l.SetOutput(logWriter)
		}
	} else {
		log.Printf("(%s) failed to create rotatelogs: %s", logger.FileName, err)
	}

	logger.Instance = l

	return logger
}

func (logger *Logger) reset() {
	logger.FileName = ""
	logger.Multi = false
	logger.Format = nil
	logger.Level = 0
}
