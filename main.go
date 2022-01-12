package main

import (
	"fmt"

	"github.com/tanaypatankar/go-workspace/dbtests/database"
	"github.com/tanaypatankar/go-workspace/dbtests/routes"
)

func main() {
	fmt.Println("hi")
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
