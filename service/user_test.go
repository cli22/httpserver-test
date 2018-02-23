package service

import (
	"testing"
	"httpserver-test/dao"
)

type userCase struct {
	name string
	user *dao.User
}

var users = []userCase{
	{"Alice", &dao.User{Id: 1, Name: "Alice", Type: "user"}},
	{"Tom", &dao.User{Id: 2, Name: "Tom", Type: "user"}},
}

func TestCreateUser(t *testing.T) {
	for _, user := range users {
		v, _ := CreateUser(user.name)
		if v.Name != user.user.Name || v.Type != user.user.Type {
			t.Error(
				"For", user.name,
				"expected", user.user,
				"got", v,
			)
		}
	}
}
