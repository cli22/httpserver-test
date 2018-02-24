package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	srv "httpserver-test/service"
	"httpserver-test/dao"
	"httpserver-test/error"
	"httpserver-test/log"
)

func GetRelationshipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	i, err := strconv.Atoi(userId)
	if err != nil {
		log.Warning.Println("GetRelationshipHandler Atoi user_id error: ", err)
		// todo return {"error":errorMsg}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := srv.ListUserRelationship(i)
	if err != nil {
		log.Warning.Println("GetRelationshipHandler ListUserRelationship error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		log.Warning.Println("GetRelationshipHandler Marshal error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info.Println("GetRelationshipHandler ListUserRelationship success, result: ", res)

	// Todo interceptor
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func CreateRelationshipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	uid, err := strconv.Atoi(userId)
	if err != nil {
		log.Warning.Println("CreateRelationshipHandler Atoi user_id error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// todo parameter check user_id, other_user_id
	otherUid := vars["other_user_id"]
	ouid, err := strconv.Atoi(otherUid)
	if err != nil {
		log.Warning.Println("CreateRelationshipHandler Atoi otherUid error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var input map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Warning.Println("CreateRelationshipHandler Unmarshal input error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := input["state"].(string)
	if state != dao.Liked && state != dao.Disliked {
		log.Warning.Println("CreateRelationshipHandler error: ", error.ErrStateInvalid)
		http.Error(w, error.ErrStateInvalid.Error(), http.StatusBadRequest)
		return
	}

	log.Info.Println("CreateRelationshipHandler parameter", uid, ouid, state)

	res, err := srv.UpdateRelationship(uid, ouid, state)
	if err != nil {
		log.Warning.Println("CreateRelationshipHandler UpdateRelationship error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info.Println("CreateRelationshipHandler success, result: ", res)

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		log.Warning.Println("CreateRelationshipHandler Unmarshal result error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Todo interceptor
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
