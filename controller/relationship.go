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

type response struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func writeResponse(w http.ResponseWriter, rsp response) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(rsp)
}

func GetRelationshipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	uid, err := strconv.Atoi(userId)
	if err != nil {
		log.Warning.Println("GetRelationshipHandler Atoi user_id error: ", err)
		writeResponse(w, response{Errno: int(error.ErrGetRelationship), Errmsg: error.Msg[error.ErrGetRelationship],})
		return
	}

	// check uid whether exists
	exist, err := srv.IsUserExist(uid)
	if err != nil {
		log.Warning.Println("GetRelationshipHandler IsUserExist user_id error: ", err)
		writeResponse(w, response{Errno: int(error.ErrGetRelationship), Errmsg: error.Msg[error.ErrGetRelationship],})
		return
	}

	if !exist {
		log.Warning.Println("GetRelationshipHandler parameter error: ", error.Msg[error.ErrUidNotExist])
		writeResponse(w, response{Errno: int(error.ErrUidNotExist), Errmsg: error.Msg[error.ErrUidNotExist],})
		return
	}

	res, err := srv.GetUserRelationship(uid)
	if err != nil {
		log.Warning.Println("GetRelationshipHandler GetUserRelationship error: ", err)
		writeResponse(w, response{Errno: int(error.ErrGetRelationship), Errmsg: error.Msg[error.ErrGetRelationship],})
		return
	}

	log.Info.Println("GetRelationshipHandler ListUserRelationship success, result: ", res)

	writeResponse(w, response{Errno: int(error.ErrOk), Errmsg: error.Msg[error.ErrOk], Data: res})
}

func UpdateRelationshipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	uid, err := strconv.Atoi(userId)
	if err != nil {
		log.Warning.Println("UpdateRelationshipHandler Atoi user_id error: ", err)
		writeResponse(w, response{Errno: int(error.ErrPutRelationship), Errmsg: error.Msg[error.ErrPutRelationship],})
		return
	}

	// check uid whether exists
	exist, err := srv.IsUserExist(uid)
	if err != nil {
		log.Warning.Println("UpdateRelationshipHandler IsUserExist user_id error: ", err)
		writeResponse(w, response{Errno: int(error.ErrPutRelationship), Errmsg: error.Msg[error.ErrPutRelationship],})
		return
	}

	if !exist {
		log.Warning.Println("UpdateRelationshipHandler parameter error: ", error.Msg[error.ErrUidNotExist])
		writeResponse(w, response{Errno: int(error.ErrUidNotExist), Errmsg: error.Msg[error.ErrUidNotExist],})
		return
	}

	// check other_user_id whether exists
	otherUid := vars["other_user_id"]
	ouid, err := strconv.Atoi(otherUid)
	if err != nil {
		log.Warning.Println("UpdateRelationshipHandler Atoi otherUid error: ", err)
		writeResponse(w, response{Errno: int(error.ErrPutRelationship), Errmsg: error.Msg[error.ErrPutRelationship],})
		return
	}

	exist, err = srv.IsUserExist(ouid)
	if err != nil {
		log.Warning.Println("UpdateRelationshipHandler IsUserExist other_user_id error: ", err)
		writeResponse(w, response{Errno: int(error.ErrPutRelationship), Errmsg: error.Msg[error.ErrPutRelationship],})
		return
	}

	if !exist {
		log.Warning.Println("UpdateRelationshipHandler parameter error: ", error.Msg[error.ErrOtherUidNotExist])
		writeResponse(w, response{Errno: int(error.ErrOtherUidNotExist), Errmsg: error.Msg[error.ErrOtherUidNotExist],})
		return
	}

	var input map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Warning.Println("UpdateRelationshipHandler Unmarshal input error: ", err)
		writeResponse(w, response{Errno: int(error.ErrPutRelationship), Errmsg: error.Msg[error.ErrPutRelationship],})
		return
	}

	state := input["state"].(string)
	if state != dao.Liked && state != dao.Disliked {
		log.Warning.Println("UpdateRelationshipHandler error: ", error.Msg[error.ErrStateInvalid])
		writeResponse(w, response{Errno: int(error.ErrStateInvalid), Errmsg: error.Msg[error.ErrStateInvalid],})
		return
	}

	log.Info.Println("UpdateRelationshipHandler parameter", uid, ouid, state)

	res, err := srv.UpdateRelationship(uid, ouid, state)
	if err != nil {
		log.Warning.Println("UpdateRelationshipHandler UpdateRelationship error: ", err)
		writeResponse(w, response{Errno: int(error.ErrPutRelationship), Errmsg: error.Msg[error.ErrPutRelationship],})
		return
	}

	log.Info.Println("UpdateRelationshipHandler success, result: ", res)

	writeResponse(w, response{int(error.ErrOk), error.Msg[error.ErrOk], res})
}
