package app

// todo which log
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/httpserver-test/app/controller"
	"github.com/httpserver-test/app/dao"
	"github.com/httpserver-test/config"
	slog "github.com/httpserver-test/log"
	srv "github.com/httpserver-test/app/service"
)

func initServer() {
	// init config
	if err := config.ParseConfig(); err != nil {
		log.Fatalf("parse config error %v", err)
	}

	// init log
	if err := slog.InitLog(config.Conf); err != nil {
		log.Fatalf("init log error %v", err)
	}

	// init db
	dao.Db = dao.NewPg(config.Conf)
	slog.Info.Println("init pg success")

	// init user, relationship service
	srv.UserSvc = srv.NewUser()
	srv.RelationshipSvc = srv.NewRelationship()
}

func Start() {
	// init
	initServer()

	router := mux.NewRouter().StrictSlash(true)
	// users
	router.HandleFunc("/users", controller.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", controller.CreateUserHandler).Methods("POST")
	// relationships
	router.HandleFunc("/users/{user_id:[0-9]+}/relationships", controller.GetRelationshipHandler).Methods("GET")
	router.HandleFunc("/users/{user_id:[0-9]+}/relationships/{other_user_id:[0-9]+}", controller.UpdateRelationshipHandler).Methods("PUT")

	// Middleware for post and put
	controller.Mw = controller.NewMw()
	router.Use(controller.Mw.MiddlewareFunc)

	slog.Error.Println(http.ListenAndServe(config.Conf.Server.Port, router))
}
