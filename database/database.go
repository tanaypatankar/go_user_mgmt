package database

import (
	"github.com/tanaypatankar/go-workspace/dbtests/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DB_Init() {
	d, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/gotables?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = d
	DB.AutoMigrate(&models.User{})
}
