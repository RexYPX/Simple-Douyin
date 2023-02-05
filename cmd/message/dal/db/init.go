package db

import (
	"Simple-Douyin/pkg/constants"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	klog.Debug("Init()::DB: ", DB)

	DB.AutoMigrate(&Message{})

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
}