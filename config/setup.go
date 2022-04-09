package config

import (
	"fmt"
	"rest/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

const DB_USERNAME = "root"
const DB_PASSWORD = "@h1ANKITBORA"
const DB_NAME = "gorm_db2"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func InitDb() *gorm.DB {
	Db = connectDB()

	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&entity.User{}, &entity.Address{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
