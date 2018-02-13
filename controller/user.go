package controller

import (
	"net/http"

	srv "httpserver-test/service"
	"fmt"
)

// 参数校验
func UserHandler(rsp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res, err := srv.ListUser()
		if err != nil{
			fmt.Print(err)
		}
		fmt.Print(res)
	case "POST":
		res, err := srv.CreateUser()
		if err != nil{
			fmt.Print(err)
		}
		fmt.Print(res)
	}
}
