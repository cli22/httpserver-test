package service

import (
	"httpserver-test/dao"
	"httpserver-test/log"
)

func ListUser() (users []dao.User, err error) {
	err = dao.Db.Model(&users).Order("id").Select()
	//err = dao.Db.Model(&users).OrderExpr("user.id ASC").Select()
	if err != nil {
		log.Warning.Println("ListUser SELECT error: ", err)
	}

	return
}

func CreateUser(name string) (user *dao.User, err error) {
	user = &dao.User{
		Name: name,
		Type: "user",
	}

	err = dao.Db.Insert(user)
	if err != nil {
		log.Warning.Println("CreateUser INSERT error: ", err)
	}

	return
}
