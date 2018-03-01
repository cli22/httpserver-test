package controller

import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	srv "httpserver-test/service"
	"httpserver-test/error"
	"httpserver-test/log"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	res, err := srv.GetUser()
	if err != nil {
		log.Warning.Println("UserHandler GetUser error: ", err)
		writeResponse(w, response{error.ErrGetUser, error.Msg[error.ErrGetUser], ""})
		return
	}

	log.Info.Println("UserHandler GetUser success, result: ", res)

	writeResponse(w, response{error.ErrOk, error.Msg[error.ErrOk], res})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// todo name need to be unique?
	var input map[string]interface{}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warning.Println("UserHandler ReadAll Body error: ", err)
		writeResponse(w, response{error.ErrCreateUser, error.Msg[error.ErrCreateUser], ""})
		return
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Warning.Println("UserHandler Unmarshal input error: ", err)
		writeResponse(w, response{error.ErrCreateUser, error.Msg[error.ErrCreateUser], ""})
		return
	}

	name := input["name"].(string)
	if name == "" {
		log.Warning.Println("UserHandler error: ", error.Msg[error.ErrNameEmpty])
		writeResponse(w, response{error.ErrNameEmpty, error.Msg[error.ErrNameEmpty], ""})
		return
	}

	res, err := srv.CreateUser(name)
	if err != nil {
		log.Warning.Println("UserHandler CreateUser error: ", err)
		writeResponse(w, response{error.ErrCreateUser, error.Msg[error.ErrCreateUser], ""})
		return
	}

	log.Info.Println("UserHandler CreateUser success, result: ", res)

	writeResponse(w, response{error.ErrOk, error.Msg[error.ErrOk], res})
}
