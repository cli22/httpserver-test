package service

import (
	"httpserver-test/app/dao"
	"httpserver-test/log"
	"httpserver-test/app/entity"
)

var User_svc *User

type User struct {
	my_dao_user *dao.MyUser
}

func NewUser() *User {
	user := new(User)
	user.my_dao_user = dao.NewMyUser()
	return user
}

func (u *User) GetUser() (users []*entity.User, err error) {
	users, err = u.my_dao_user.List()
	if err != nil {
		log.Warning.Println("GetUser error: ", err)
	}

	return
}

func (u *User) CreateUser(data *entity.User) (users []*entity.User, err error) {
	users, err = u.my_dao_user.Add(data)
	if err != nil {
		log.Warning.Println("CreateUser error: ", err)
	}

	return
}

func (u *User) IsUserExist(data *entity.User) (exist bool, err error) {
	exist, err = u.my_dao_user.IsUserExist(data)
	if err != nil {
		log.Warning.Println("IsUserExist error: ", err)
	}

	return
}
