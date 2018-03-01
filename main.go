package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"httpserver-test/controller"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// users
	router.HandleFunc("/users", controller.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", controller.CreateUserHandler).Methods("POST")
	// relationships
	router.HandleFunc("/users/{user_id:[0-9]+}/relationships", controller.GetRelationshipHandler).Methods("GET")
	router.HandleFunc("/users/{user_id:[0-9]+}/relationships/{other_user_id:[0-9]+}", controller.UpdateRelationshipHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
