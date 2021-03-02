package logger

import (
	"gin-api/internal/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"path"
	"time"
)

const (
	TimestampFormat = "2006-01-02 15:04:05"
)

type Fields logrus.Fields

func (fields Fields) Fields() logrus.Fields {
	return logrus.Fields(fields)
}

type Logger struct {
	Instance	*logrus.Logger
	Formatter   logrus.Formatter
	Level       logrus.Level
	OpenLFSHOOK bool
}

// 实例化 Logger
func New() *Logger {
	return &Logger{
		Instance: logrus.New(),
	}
}

// 获取日志实例
func (logger *Logger) GetInstance() *logrus.Logger {
	return logger.Instance
}

// 获取新的日志实例
func NewInstance() *logrus.Logger {
	return New().GetInstance()
}

// 获取新的日志写入实例
func NewWriteInstance(filename string) *logrus.Logger {
	return NewWrite(filename).GetInstance()
}

// 日志写入
func NewWrite(filename string) *Logger {
	// 实例化
	logger := New()

	// 设置日志级别
	if logger.Level == 0 {
		logger.Instance.SetLevel(logrus.DebugLevel)
	} else {
		logger.Instance.SetLevel(logger.Level)
	}

	// 设置日志格式
	if logger.Formatter == nil {
		logger.Instance.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: TimestampFormat,
		})
	} else {
		logger.Instance.SetFormatter(logger.Formatter)
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
		logger.Instance.SetOutput(logWriter)
	} else {
		log.Printf("(%s) failed to create rotatelogs: %s", filename, err)
	}

	// 开启 lfshook
	if logger.OpenLFSHOOK {
		lfs := open_lfshook(logWriter)
		// 新增 Hook
		logger.Instance.AddHook(lfs)
	}

	return logger
}

// 开启 lfshook
func open_lfshook(logWriter *rotatelogs.RotateLogs) *lfshook.LfsHook {
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	//设置日志格式
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: TimestampFormat,
	})

	return lfHook
}
