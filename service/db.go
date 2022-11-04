package service

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init() {
	var err error
	if db, err = gorm.Open(postgres.Open(
		"host=127.0.0.1 user=postgres password=root dbname=video-server port=5432 sslmode=disable TimeZone=Asia/Shanghai"+
			" sslmode=disable TimeZone=Asia/Shanghai"),
		&gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}}); err != nil {
		zap.L().Fatal("connect database failed", zap.Error(err))
	}

	db.Logger = db.Logger.LogMode(logger.Silent)
}

func DB() *gorm.DB {
	return db
}
