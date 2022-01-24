package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tanaypatankar/go_user_mgmt/controller"
)

func InitializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users/", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controller.GetUserByID).Methods("GET")
	r.HandleFunc("/users/", controller.CreateUser).Methods("POST")
	r.HandleFunc("/users/", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", nil))
}
