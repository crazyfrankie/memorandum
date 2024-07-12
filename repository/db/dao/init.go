package dao

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"memorandum/config"
	"memorandum/repository/db/model"
)

var (
	DB *gorm.DB
)

func InitDB() {
	dsn := strings.Join([]string{config.User, ":", config.Password, "@tcp(", config.Host, ":", config.Port, ")/", config.Db, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: ormLogger, // 打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加s
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.User{}, &model.Task{})

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池
	sqlDB.SetMaxOpenConns(100) // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
}
