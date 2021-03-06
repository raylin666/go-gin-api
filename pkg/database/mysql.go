package database

import (
	"fmt"
	"gin-api/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

var database map[string]*gorm.DB

func InitDatabase() {
	var (
		err  error
		conn *gorm.DB
	)

	database = make(map[string]*gorm.DB)

	c := config.Get().Database

	for key, value := range c {
		var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			value.UserName,
			value.Password,
			value.Host,
			value.Port,
			value.DbName,
			value.Charset)

		conn, err = gorm.Open(
			mysql.Open(dsn),
			&gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   value.Prefix,	// 设置表前缀
					SingularTable: true,			// 全局禁用表名复数
				},
			})

		if err != nil {
			log.Panic(err)
		}

		db, _ := conn.DB()
		// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
		db.SetMaxIdleConns(value.MaxIdleConn)
		// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
		db.SetMaxOpenConns(value.MaxOpenConn)
		// 设置最大连接超时
		db.SetConnMaxLifetime(time.Minute * value.MaxLifeTime)

		// 使用插件
		_ = conn.Use(&TracePlugin{})

		database[strings.ToLower(key)] = conn
	}
}

func GetDB(connection string) *gorm.DB {
	return database[connection]
}
