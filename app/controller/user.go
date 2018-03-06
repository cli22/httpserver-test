package controller

import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	"httpserver-test/app/error"
	"httpserver-test/log"
	srv "httpserver-test/app/service"
	"httpserver-test/app/entity"
)

type Middleware struct {
	input map[string]interface{}
}

func (mw *Middleware) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "PUT" {
			defer r.Body.Close()
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Warning.Println("UserHandler error: ", err)
				writeResponse(w, response{Errno: error.ErrCreateUser, Errmsg: error.Msg[error.ErrCreateUser]})
				return
			}

			//todo
			err = json.Unmarshal(body, &mw.input)
			if err != nil {
				log.Warning.Println("UserHandler error: ", err)
				writeResponse(w, response{Errno: error.ErrCreateUser, Errmsg: error.Msg[error.ErrCreateUser]})
				return
			}
		}
		log.Info.Println("mw.input", mw.input)
		next.ServeHTTP(w, r)
	})
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	res, err := srv.User_svc.GetUser()

	if err != nil {
		log.Warning.Println("UserHandler GetUser error: ", err)
		writeResponse(w, response{Errno: error.ErrGetUser, Errmsg: error.Msg[error.ErrGetUser]})
		return
	}

	log.Info.Println("UserHandler GetUser result: ", res)

	writeResponse(w, response{Errno: error.ErrOk, Errmsg: error.Msg[error.ErrOk], Data: res})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input map[string]interface{}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	//todo
	_ = json.Unmarshal(body, &input)

	name := input["name"].(string)
	if name == "" {
		log.Warning.Println("UserHandler error: ", error.Msg[error.ErrNameEmpty])
		writeResponse(w, response{Errno: error.ErrNameEmpty, Errmsg: error.Msg[error.ErrNameEmpty]})
		return
	}

	user := new(entity.User)
	user.Name = name

	res, err := srv.User_svc.CreateUser(user)
	if err != nil {
		log.Warning.Println("CreateUser error: ", err)
		writeResponse(w, response{Errno: error.ErrCreateUser, Errmsg: error.Msg[error.ErrCreateUser]})
		return
	}

	log.Info.Println("CreateUser result: ", res)

	writeResponse(w, response{Errno: error.ErrOk, Errmsg: error.Msg[error.ErrOk], Data: res})
}
