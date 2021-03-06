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
	c *conf
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
	// 文件名称(文件写入时存在值)
	FileName string
	// 配置信息
	Conf     *conf
}

type conf struct {
	level  logrus.Level
	format logrus.Formatter
	// 是否并发写入文件及控制台打印
	multi bool
}

func InitLogger() {
	// 创建文件夹
	utils.CreateDirectory(config.Get().Log.Path)

	c = &conf{}

	// 注册日志写实例
	register()
}

func register() map[string]*Logger {
	loggerWriteMaps = map[string]*Logger{
		constant.LOG_MULTI_APP: instanceMulti(constant.LOG_MULTI_APP),
		constant.LOG_APP:       instance(constant.LOG_APP),
		constant.LOG_DB:        instance(constant.LOG_DB),
		constant.LOG_REDIS:     instance(constant.LOG_REDIS),
		constant.LOG_REQUEST:   instance(constant.LOG_REQUEST),
		constant.LOG_SQL:   	instance(constant.LOG_SQL),
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

func instance(filename string) *Logger {
	return c.create(filename)
}

func instanceMulti(filename string) *Logger {
	c.multi = true
	return c.create(filename)
}

func (c *conf) reset()  {
	c.multi = false
	c.format = nil
	c.level = 0
}

// 创建 Logger 实例 初始化配置
func (c *conf) create(filename string) *Logger {
	logger := logrus.New()

	// 设置日志级别
	if c.level == 0 {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(c.level)
	}

	// 设置日志格式
	if c.format == nil {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: TimestampFormat,
		})
	} else {
		logger.SetFormatter(c.format)
	}

	file := path.Join(config.Get().Log.Path, filename)
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
		if c.multi {
			logger.SetOutput(io.MultiWriter(os.Stdout, logWriter))
		} else {
			logger.SetOutput(logWriter)
		}
	} else {
		log.Printf("(%s) failed to create rotatelogs: %s", filename, err)
	}

	new_logger := &Logger{
		Instance: logger,
		FileName: filename,
		Conf:     c,
	}

	c.reset()

	return new_logger
}
