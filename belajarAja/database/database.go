package database

import (
	"belajaraja/entity"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB gorm.DB

func Database() *gorm.DB {
	dsn := "rayhan:rayhan222@tcp(127.0.0.1:3306)/golangDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("cannot connect to databases")
	}

	fmt.Println("sucsess connect to db!")

	DB = *db
	DB.AutoMigrate(&entity.User{}, &entity.Ujian{})
	return db
}
