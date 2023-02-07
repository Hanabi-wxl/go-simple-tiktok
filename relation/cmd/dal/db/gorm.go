package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"relation/pkg/consts"
	"time"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(consts.MysqlDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 默认不加复数
		}})
	if err != nil {
		log.Fatal("数据库链接失败: ", err)
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
}
