package controller

import (
	"net/http"

	srv "httpserver-test/service"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res, err := srv.ListUser()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(res)
		jsonBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Print(err)
		}
		// Todo interceptor
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)

	case "POST":
		// todo parameter check
		var user map[string]interface{}
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &user)
		res, err := srv.CreateUser(user["name"].(string))
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(res)
		jsonBytes, err := json.Marshal(res)
		if err != nil {
			fmt.Print(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}
}
