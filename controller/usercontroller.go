package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tanaypatankar/go_user_mgmt/database"
	"github.com/tanaypatankar/go_user_mgmt/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// fmt.Println(mux.Vars(r))
	var users []models.User
	database.DB.Find(&users)
	// fmt.Printf("%T", users[0])
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)
	defer r.Body.Close()
	var user models.User

	if result := database.DB.First(&user, id); result.Error != nil {
		log.Println(result.Error)
		err := models.Error{
			Message: result.Error.Error(),
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)

	} else {
		json.NewEncoder(w).Encode(user)
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if result := database.DB.Create(&user); result.Error != nil {
		log.Println(result.Error)
		err := models.Error{
			Message: result.Error.Error(),
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	var original_user models.User
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&user)
	// Checks if record exists
	if result := database.DB.First(&original_user, user.Id); result.Error != nil {
		log.Println(result.Error)
		err := models.Error{
			Message: result.Error.Error(),
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
	} else {

		database.DB.Save(&user)
		json.NewEncoder(w).Encode(user)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	var user models.User
	vars := mux.Vars(r)
	defer r.Body.Close()

	if result := database.DB.First(&user, vars["id"]); result.Error != nil {
		log.Println(result.Error)
		err := models.Error{
			Message: result.Error.Error(),
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
	} else {
		database.DB.Unscoped().Delete(&user)
		json.NewEncoder(w).Encode(user)
	}

}
