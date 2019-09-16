package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init(conn string) {
	fmt.Println(conn)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
	}

	db.Debug()
	db.LogMode(true)
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)

	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Video{})
}
