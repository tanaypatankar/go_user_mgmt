package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tanaypatankar/go_user_mgmt/database"
	"github.com/tanaypatankar/go_user_mgmt/routes"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("hi")
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	database.DB_Init()
	// user := models.User{
	// 	Id:      1,
	// 	Name:    "Tanay",
	// 	Email:   "tanay@platform9.com",
	// 	Age:     21,
	// 	Gender:  "M",
	// 	Country: "India",
	// 	Status:  "Active",
	// }
	// result := database.DB.Create(&user)
	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)

	routes.InitializeRouter()

}
