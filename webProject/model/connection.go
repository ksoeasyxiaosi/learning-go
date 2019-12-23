package model

import (
	"github.com/jinzhu/gorm"
	"os"
	"time"

	// 必须要导入 否则数据库连接出报错
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitConnection(connectionString string) {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("数据库连接失败")
	}

	if os.Getenv("GIN_MODE") != "product" {
		db.LogMode(true)
	}

	// 设置连接池
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
}
