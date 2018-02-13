package controller

import (
	"net/http"

	srv "httpserver-test/service"

	"fmt"
)

func GetRelationshipHandler(rsp http.ResponseWriter, req *http.Request) {
	res, err := srv.ListUserRelationship(11)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(res)
}

func CreateRelationshipHandler(rsp http.ResponseWriter, req *http.Request) {
	res, err := srv.UpdateRelationship(11, 12, srv.Liked)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(res)
}
