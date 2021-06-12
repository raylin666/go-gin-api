package database

import (
	"github.com/raylin666/go-utils/database"
	"github.com/raylin666/go-utils/logger"
	"go-gin-api/config"
	"go-gin-api/internal/constant"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"time"
)

func InitDatabase() {
	var (
		c  map[string]*database.DatabaseConfig
		cr = config.Get().Database
	)

	c = make(map[string]*database.DatabaseConfig, len(cr))

	for key, value := range cr {
		rc := &database.DatabaseConfig{
			Driver:      value.Driver,
			DbName:      value.DbName,
			Host:        value.Host,
			UserName:    value.UserName,
			Password:    value.Password,
			Charset:     value.Charset,
			Port:        value.Port,
			Prefix:      value.Prefix,
			MaxIdleConn: value.MaxIdleConn,
			MaxOpenConn: value.MaxOpenConn,
			MaxLifeTime: value.MaxLifeTime,
			ParseTime:   value.ParseTime,
			Loc:         value.Loc,
			OpenPlugin:  value.OpenPlugin,
		}
		c[key] = rc
	}

	database.InitDatabase(c, &database.PluginConfig{
		After: func(gormDb *gorm.DB, sql string, ts time.Time) {
			logger.NewWrite(constant.LogSql).WithFields(logger.H{
				"query":       sql,
				"rows":        gormDb.Statement.RowsAffected,
				"stack":       utils.FileWithLineNum(),
				"costSeconds": time.Since(ts).Seconds(),
			}.Fields()).Info()
		},
	})
}

// 获取链接
func Get(connection string) *database.Database {
	return database.Get(connection)
}

// 关闭链接
func Close(connection string) error {
	return database.Close(connection)
}

// 关闭所有链接
func CloseAll() error {
	return database.CloseAll()
}
