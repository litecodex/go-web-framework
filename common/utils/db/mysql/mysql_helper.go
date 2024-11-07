package mysql

import (
	"fmt"
	DBLogger "github.com/litecodex/go-web-framework/common/utils/db/logger"
	DBModel "github.com/litecodex/go-web-framework/common/utils/db/model"
	"github.com/litecodex/go-web-framework/web/utils/logger"
	"go.uber.org/zap"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(config *DBModel.DBConfig, log *zap.Logger) *gorm.DB {
	if config.Username == "" {
		panic("请输入登录用户名，字段：config.Username")
	}
	if config.Password == "" {
		panic("请输入登录密码，字段：config.Password")
	}
	if config.Host == "" {
		panic("请输入ip地址，字段：config.Host")
	}
	if config.Port == "" {
		panic("请输入数据库端口，字段：config.Port")
	}
	if config.DBName == "" {
		panic("请输入数据库名，字段：config.DBName")
	}

	var dbLogger *DBLogger.GormLogger
	if log != nil {
		dbLogger = DBLogger.NewGormLogger(log)
	} else {
		dbLogger = DBLogger.NewGormLogger(logger.NewConsoleLogger(0))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.DBName)
	db, err := gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}
