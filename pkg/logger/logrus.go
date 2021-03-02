package logger

import (
	"fmt"
	"gin-api/internal/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"time"
)

const (
	TimestampFormat = "2006-01-02 15:04:05"
	// 默认写入的文件名称
	DefaultFileName = "app"
)

var (
	// 日志实例集合
	loggerInstanceMaps map[string]*Logger
)

type Fields logrus.Fields

func (fields Fields) Fields() logrus.Fields {
	return logrus.Fields(fields)
}

type Logger struct {
	// 日志实例
	Instance    *logrus.Logger
	// 日志格式
	Formatter   logrus.Formatter
	// 日志级别
	Level       logrus.Level
	// 是否并发写入文件及控制台打印
	IsMulti     bool
	// 是否开启 lfshook
	OpenLFSHOOK bool
}

func InitLogger() {
	// 创建文件夹
	createDirectory()
}

// 获取日志新实例
func New() *logrus.Logger {
	return logrus.New()
}

func (logger Logger) Write(filename string) *logrus.Logger {
	return write(filename, &logger)
}

// 获取日志实例并写入文件
func NewWrite(filename string) *logrus.Logger {
	return write(filename, &Logger{})
}

// 获取日志实例并写入文件及打印到控制台
func NewMuWrite(filename string) *logrus.Logger {
	return write(filename, &Logger{
		IsMulti: true,
	})
}

func write(filename string, l *Logger) *logrus.Logger {
	var (
		logger *Logger
		ok     bool
	)

	if loggerInstanceMaps == nil {
		loggerInstanceMaps = make(map[string]*Logger)
	}

	if logger, ok = loggerInstanceMaps[filename]; !ok {
		newLogger := logrus.New()

		// 设置日志级别
		if l.Level == 0 {
			newLogger.SetLevel(logrus.DebugLevel)
		} else {
			newLogger.SetLevel(l.Level)
		}

		// 设置日志格式
		if l.Formatter == nil {
			newLogger.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: TimestampFormat,
			})
		} else {
			newLogger.SetFormatter(l.Formatter)
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
			if l.IsMulti {
				newLogger.SetOutput(io.MultiWriter(os.Stdout, logWriter))
			} else {
				newLogger.SetOutput(logWriter)
			}
		} else {
			log.Printf("(%s) failed to create rotatelogs: %s", filename, err)
		}

		// 开启 lfshook
		if l.OpenLFSHOOK {
			lfs := openLfshook(logWriter)
			// 新增 Hook
			newLogger.AddHook(lfs)
		}

		// 设置 Logger
		logger = &Logger{
			Instance:    newLogger,
			Formatter:   l.Formatter,
			Level:       l.Level,
			IsMulti:     l.IsMulti,
			OpenLFSHOOK: l.OpenLFSHOOK,
		}

		loggerInstanceMaps[filename] = logger
	}

	return logger.Instance
}

// 判断日志文件夹是否存在,不存在则创建
func createDirectory() {
	dir := config.Get().Log.Path
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				panic(fmt.Errorf("Fatal create %v directory %v \n", dir, err))
			} else {
				fmt.Printf("create %v directory success \n", dir)
			}
		}
	}
}

// 开启 lfshook
func openLfshook(logWriter *rotatelogs.RotateLogs) *lfshook.LfsHook {
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
