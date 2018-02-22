package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	srv "httpserver-test/service"
	"httpserver-test/dao"

	"github.com/gorilla/mux"
	"encoding/json"
)

func GetRelationshipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	i, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	res, err := srv.ListUserRelationship(i)
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
}

func CreateRelationshipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	uid, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println(err)
	}
	otherUid := vars["other_user_id"]
	ouid, err := strconv.Atoi(otherUid)
	if err != nil {
		fmt.Println(err)
	}
	// todo parameter check user_id, other_user_id
	var state map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &state)
	s := state["state"].(string)
	if s != string(dao.Liked) && s != string(dao.Disliked) {
		fmt.Println("parameter invaild: state")
	}
	fmt.Println(uid, ouid, s)
	res, err := srv.UpdateRelationship(uid, ouid, s)
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
}
