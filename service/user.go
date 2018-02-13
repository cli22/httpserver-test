package service

import (
	"fmt"
	//"log" //Todo

	"httpserver-test/dao"
)

func ListUser() (users []dao.User, err error) {
	err = dao.Db.Model(&users).Select()
	if err != nil {
		fmt.Print(err)
	}
	return
}

func CreateUser() (user dao.User, err error) {
	err = dao.Db.Insert(user)
	if err != nil {
		fmt.Print(err)
	}
	return
}
