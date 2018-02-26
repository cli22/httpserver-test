package service

import (
	"httpserver-test/dao"
	"httpserver-test/log"
)

func GetUser() (users []dao.User, err error) {
	err = dao.Db.Model(&users).Order("id ASC").Select()
	if err != nil {
		log.Warning.Println("GetUser SELECT db error: ", err)
	}

	return
}

func CreateUser(name string) (user *dao.User, err error) {
	user = &dao.User{
		Name: name,
		Type: dao.UserType,
	}

	err = dao.Db.Insert(user)
	if err != nil {
		log.Warning.Println("CreateUser INSERT db error: ", err)
	}

	return
}

func IsUserExist(uid int) (exist bool, err error) {
	count, err := dao.Db.Model(&dao.User{}).Where("id=?", uid).Count()
	if err != nil {
		log.Warning.Println("IsUserExists SELECT db error: ", err)
	}

	if count > 0 {
		exist = true
	} else {
		exist = false
	}

	return
}
