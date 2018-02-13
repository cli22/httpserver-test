package main

import (
	"log"
	"net/http"

	"httpserver-test/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// users
	router.HandleFunc("/users", controller.UserHandler).Methods("GET", "POST")
	// relationships
	router.HandleFunc("/users/:user_id/relationships", controller.GetRelationshipHandler).Methods("GET")
	router.HandleFunc("/users/:user_id/relationships/:other_user_id", controller.CreateRelationshipHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":80", router))
}
