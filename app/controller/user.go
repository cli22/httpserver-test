package controller

import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/httpserver-test/app/entity"
	"github.com/httpserver-test/app/error"
	"github.com/httpserver-test/log"
	srv "github.com/httpserver-test/app/service"
)

var Mw *Middleware

type Middleware struct {
	input map[string]interface{}
}

func NewMw() *Middleware {
	return new(Middleware)
}

// for put and post method, unmarshal body content for controller handle use.
// todo add parameter check for user_id and other_user_id
func (mw *Middleware) MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" || r.Method == "PUT" {
			defer r.Body.Close()
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Warning.Println("MiddlewareFunc error: ", err)
				writeResponse(w, response{Errno: error.ErrCreateUser, Errmsg: error.Msg[error.ErrCreateUser]})
				return
			}

			//todo
			err = json.Unmarshal(body, &mw.input)
			if err != nil {
				log.Warning.Println("MiddlewareFunc error: ", err)
				writeResponse(w, response{Errno: error.ErrCreateUser, Errmsg: error.Msg[error.ErrCreateUser]})
				return
			}
		}
		log.Info.Println("mw.input", mw.input)
		next.ServeHTTP(w, r)
	})
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	res, err := srv.UserSvc.GetUser()

	if err != nil {
		log.Warning.Println("UserHandler GetUser error: ", err)
		writeResponse(w, response{Errno: error.ErrGetUser, Errmsg: error.Msg[error.ErrGetUser]})
		return
	}

	log.Info.Println("UserHandler GetUser result: ", res)

	writeResponse(w, response{Errno: error.ErrOk, Errmsg: error.Msg[error.ErrOk], Data: res})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	name := Mw.input["name"].(string)
	if name == "" {
		log.Warning.Println("UserHandler error: ", error.Msg[error.ErrNameEmpty])
		writeResponse(w, response{Errno: error.ErrNameEmpty, Errmsg: error.Msg[error.ErrNameEmpty]})
		return
	}

	user := new(entity.User)
	user.Name = name

	res, err := srv.UserSvc.CreateUser(user)
	if err != nil {
		log.Warning.Println("CreateUser error: ", err)
		writeResponse(w, response{Errno: error.ErrCreateUser, Errmsg: error.Msg[error.ErrCreateUser]})
		return
	}

	log.Info.Println("CreateUser result: ", res)

	writeResponse(w, response{Errno: error.ErrOk, Errmsg: error.Msg[error.ErrOk], Data: res})
}
