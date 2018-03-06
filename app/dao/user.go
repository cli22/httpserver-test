package dao

import (
	"strconv"

	"github.com/httpserver-test/app/entity"
	"github.com/httpserver-test/log"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type UserTypeGroup = string

const (
	UserType UserTypeGroup = "user"
)

type MyUser struct{}

func NewMyUser() *MyUser {
	return &MyUser{}
}

func (this *MyUser) resToUser(res []*User) (users []*entity.User, err error) {
	if len(res) == 0 {
		return nil, nil
	}

	for _, r := range res {
		user := new(entity.User)
		user.Id = strconv.Itoa(r.Id)
		user.Name = r.Name
		user.Type = r.Type

		users = append(users, user)
	}

	return users, nil
}

func (this *MyUser) List() (users []*entity.User, err error) {
	res := make([]*User, 0)

	err = Db.Model(&res).Order("id ASC").Select()
	if err != nil {
		log.Warning.Println("SELECT db error: ", err)
	}

	users, err = this.resToUser(res)

	return users, err
}

func (this *MyUser) Add(data *entity.User) (user *entity.User, err error) {
	daoUser := &User{
		Name: data.Name,
		Type: UserType,
	}

	err = Db.Insert(daoUser)
	if err != nil {
		log.Warning.Println("INSERT db error: ", err)
	}

	res := make([]*User, 0)
	res = append(res, daoUser)
	users, err := this.resToUser(res)

	if len(users) == 1 {
		user = users[0]
	}

	return user, err
}

func (this *MyUser) IsUserExist(data *entity.User) (exist bool, err error) {
	count, err := Db.Model(&User{}).Where("id=?", data.Id).Count()
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
