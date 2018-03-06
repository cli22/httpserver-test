package service

import (
	"httpserver-test/app/dao"
	"httpserver-test/app/entity"
	"httpserver-test/log"
)

var UserSvc *User

type User struct {
	myDaoUser *dao.MyUser
}

func NewUser() *User {
	user := new(User)
	user.myDaoUser = dao.NewMyUser()
	return user
}

func (u *User) GetUser() (users []*entity.User, err error) {
	users, err = u.myDaoUser.List()
	if err != nil {
		log.Warning.Println("GetUser error: ", err)
	}

	return
}

func (u *User) CreateUser(data *entity.User) (user *entity.User, err error) {
	user, err = u.myDaoUser.Add(data)
	if err != nil {
		log.Warning.Println("CreateUser error: ", err)
	}

	return
}

func (u *User) IsUserExist(data *entity.User) (exist bool, err error) {
	exist, err = u.myDaoUser.IsUserExist(data)
	if err != nil {
		log.Warning.Println("IsUserExist error: ", err)
	}

	return
}
