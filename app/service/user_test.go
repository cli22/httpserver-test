package service

import (
	"testing"
	"httpserver-test/app/dao"
	"httpserver-test/app/entity"
)

type userCase struct {
	name   string
	expect *entity.User
}

var users = []userCase{
	{"Alice", &entity.User{Name: "Alice", Type: dao.UserType}},
	{"Tom", &entity.User{Name: "Tom", Type: dao.UserType}},
}

func TestUser_CreateUser(t *testing.T) {
	daoUser := dao.NewMyUser()
	for _, user := range users {
		inputUser := &entity.User{Name: user.name}
		v, _ := daoUser.Add(inputUser)
		if v.Name != user.expect.Name || v.Type != user.expect.Type {
			t.Error(
				"For", user.name,
				"expected", user.expect,
				"got", v,
			)
		}
	}
}
