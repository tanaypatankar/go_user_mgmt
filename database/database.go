package database

import (
	"fmt"
	"log"
	"os"

	"github.com/tanaypatankar/go_user_mgmt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DB_Init() {
	conn_str := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_ADDRESS"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_TABLE"))
	d, err := gorm.Open(mysql.Open(conn_str), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	DB = d
	DB.AutoMigrate(&models.User{})
}
