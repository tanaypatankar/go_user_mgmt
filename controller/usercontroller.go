package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tanaypatankar/go-workspace/dbtests/database"
	"github.com/tanaypatankar/go-workspace/dbtests/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println(mux.Vars(r))
	var users []models.User
	database.DB.Find(&users)
	// fmt.Printf("%T", users[0])
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		err := models.Error{
			Message: "Record Not Found",
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)

	} else {
		json.NewEncoder(w).Encode(user)
	}

}
