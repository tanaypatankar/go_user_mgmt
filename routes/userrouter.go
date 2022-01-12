package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tanaypatankar/go-workspace/dbtests/controller"
)

func InitializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users/", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controller.GetUserByID).Methods("GET")
	// r.HandleFunc("/users/", controller.CreateUser).Methods("POST")
	// r.HandleFunc("/users/{userid}", controller.UpdateUser).Methods("PUT")
	// r.HandleFunc("/users/{userid}", controller.DeleteUser).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", nil))
}
