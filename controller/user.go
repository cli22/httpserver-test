package controller

import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	srv "httpserver-test/service"
	"httpserver-test/error"
	"httpserver-test/log"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res, err := srv.ListUser()
		if err != nil {
			log.Warning.Println("UserHandler ListUser error: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(res)
		if err != nil {
			log.Warning.Println("UserHandler ListUser Marshal result error: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Info.Println("UserHandler ListUser success, result: ", res)

		// Todo interceptor
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)

	case "POST":
		// todo name need to be unique?
		var input map[string]interface{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Warning.Println("UserHandler ReadAll Body error: ", err)
		}

		err = json.Unmarshal(body, &input)
		if err != nil {
			log.Warning.Println("UserHandler Unmarshal input error: ", err)
		}

		name := input["name"].(string)
		if name == "" {
			log.Warning.Println("UserHandler error: ", error.ErrNameEmpty)
			http.Error(w, error.ErrNameEmpty.Error(), http.StatusBadRequest)
			return
		}

		res, err := srv.CreateUser(name)
		if err != nil {
			log.Warning.Println("UserHandler CreateUser error: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(res)
		if err != nil {
			log.Warning.Println("UserHandler CreateUser Marshal result error: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Info.Println("UserHandler CreateUser success, result: ", res)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}
}
