package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tanaypatankar/go_user_mgmt/database"
	"github.com/tanaypatankar/go_user_mgmt/models"
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
		log.Println(result.Error)
		err := models.Error{
			Message: "Record Not Found",
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)

	} else {
		json.NewEncoder(w).Encode(user)
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	vars := mux.Vars(r)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	if result := database.DB.First(&user, vars["userid"]); result != nil {
		fmt.Println(result.Error)
	}
	json.Unmarshal(body, &user)
	id, err := strconv.Atoi(vars["userid"])
	user.ID = uint(id)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	vars := mux.Vars(r)
	defer r.Body.Close()

	if result := database.DB.First(&user, vars["userid"]); result != nil {
		fmt.Println(result.Error)
	}
	database.DB.Delete(&user)
	json.NewEncoder(w).Encode(user)
}
