package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tanaypatankar/go_user_mgmt/database"
	"github.com/tanaypatankar/go_user_mgmt/models"
	"github.com/tanaypatankar/go_user_mgmt/routes"
)

func main() {
	fmt.Println("hi")
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	database.DB_Init()
	user := models.User{
		Id:      1,
		Name:    "Tanay",
		Email:   "tanay@platform9.com",
		Age:     21,
		Gender:  "M",
		Country: "India",
		Status:  "Active",
	}
	result := database.DB.Create(&user)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	routes.InitializeRouter()

}
