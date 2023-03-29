package database

import (
	"fmt"
	"go-fiber-practice/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBName)
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("🚀 Connected Successfully to the Database")
	return db
}
