package core

import (
	"fmt"
	"gin_docs_server/global"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		logrus.Warn("未配置mysql，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface

	var logLevel logger.LogLevel = logger.Error
	switch global.Config.Mysql.LogLevel {
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	case "error":
		logLevel = logger.Error
	}
	mysqlLogger = logger.Default.LogMode(logLevel)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		logrus.Fatalln(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	return db
}