package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"

	"github.com/httpserver-test/app/dao"
	"github.com/httpserver-test/app/entity"
	"github.com/httpserver-test/app/error"
	"github.com/httpserver-test/log"
	srv "github.com/httpserver-test/app/service"
)

type response struct {
	Errno  codes.Code  `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func writeResponse(w http.ResponseWriter, rsp response) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(rsp)
}

func GetRelationshipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// check uid whether exists
	user := new(entity.User)
	user.Id = vars["user_id"]

	exist, err := srv.UserSvc.IsUserExist(user)
	if err != nil {
		log.Warning.Println("GetRelationshipHandler IsUserExist error: ", err)
		writeResponse(w, response{Errno: error.ErrGetRelationship, Errmsg: error.Msg[error.ErrGetRelationship],})
		return
	}

	if !exist {
		log.Warning.Println("GetRelationshipHandler error: ", error.Msg[error.ErrUidNotExist])
		writeResponse(w, response{Errno: error.ErrUidNotExist, Errmsg: error.Msg[error.ErrUidNotExist],})
		return
	}

	relationship := new(entity.Relationship)
	relationship.Uid = user.Id

	res, err := srv.RelationshipSvc.GetUserRelationship(relationship)
	if err != nil {
		log.Warning.Println("GetUserRelationship error: ", err)
		writeResponse(w, response{Errno: error.ErrGetRelationship, Errmsg: error.Msg[error.ErrGetRelationship],})
		return
	}

	log.Info.Println("GetUserRelationship result: ", res)

	writeResponse(w, response{Errno: error.ErrOk, Errmsg: error.Msg[error.ErrOk], Data: res})
}

func UpdateRelationshipHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["user_id"]
	otherUid := vars["other_user_id"]

	// check uid whether exists
	user := new(entity.User)
	user.Id = uid

	exist, err := srv.UserSvc.IsUserExist(user)
	if err != nil {
		log.Warning.Println("IsUserExist uid error: ", err)
		writeResponse(w, response{Errno: error.ErrGetRelationship, Errmsg: error.Msg[error.ErrGetRelationship],})
		return
	}

	if !exist {
		log.Warning.Println("IsUserExist other_uid error: ", error.Msg[error.ErrUidNotExist])
		writeResponse(w, response{Errno: error.ErrUidNotExist, Errmsg: error.Msg[error.ErrUidNotExist],})
		return
	}

	// check other_user_id whether exists
	user.Id = otherUid
	exist, err = srv.UserSvc.IsUserExist(user)
	if err != nil {
		log.Warning.Println("IsUserExist other_uid error: ", err)
		writeResponse(w, response{Errno: error.ErrGetRelationship, Errmsg: error.Msg[error.ErrGetRelationship],})
		return
	}

	if !exist {
		log.Warning.Println("IsUserExist other_uid error: ", error.Msg[error.ErrOtherUidNotExist])
		writeResponse(w, response{Errno: error.ErrOtherUidNotExist, Errmsg: error.Msg[error.ErrOtherUidNotExist],})
		return
	}

	state := Mw.input["state"].(string)
	if state != dao.Liked && state != dao.Disliked {
		log.Warning.Println("UpdateRelationshipHandler error: ", error.Msg[error.ErrStateInvalid])
		writeResponse(w, response{Errno: error.ErrStateInvalid, Errmsg: error.Msg[error.ErrStateInvalid],})
		return
	}

	relationship := new(entity.Relationship)
	relationship.Uid = uid
	relationship.OtherUid = otherUid
	relationship.State = state

	log.Info.Println("UpdateRelationshipHandler parameter", relationship)

	res, err := srv.RelationshipSvc.UpdateRelationship(relationship)
	if err != nil {
		log.Warning.Println("UpdateRelationship error: ", err)
		writeResponse(w, response{Errno: error.ErrPutRelationship, Errmsg: error.Msg[error.ErrPutRelationship],})
		return
	}

	log.Info.Println("UpdateRelationship result: ", res)

	writeResponse(w, response{error.ErrOk, error.Msg[error.ErrOk], res})
}
