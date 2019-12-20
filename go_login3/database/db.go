package database

import "github.com/jinzhu/gorm"

var DBInstance *gorm.DB

func InitDB() {
	//连接数据库（sqlite)
	db, err := gorm.Open("sqlite3", "./database/test.db")
	if err != nil {
		panic("连接数据库失败")
	}

	DBInstance = db
	db.LogMode(true)
}