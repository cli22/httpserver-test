package dao

import (
	"github.com/go-pg/pg"
)

// todo 改成配置文件
var Db = pg.Connect(&pg.Options{
	User:     "postgres",
	Password: "123",
	Database: "test",
})

//func init() {
//	err := createSchema(Db)
//	if err != nil {
//		fmt.Print(err)
//	}
//}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&User{}, &Relationship{}} {
		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

type User struct {
	Id int64 `json:"id"`
	//Uid  int64
	Name string `json:"name"`
	Type string `json:"type"`
}

type Relationship struct {
	Id       int64  `json:"-"`
	Uid      int64  `json:"-"`
	OtherUid int64  `json:"user_id"`
	State    string `json:"state"`
	Type     string `json:"type"`
}

type RelationshipState = string //typedef

const (
	Liked    RelationshipState = "liked"
	Disliked RelationshipState = "disliked"
	Matched  RelationshipState = "matched"
	Default  RelationshipState = ""
)

type UserTypeGroup = string

const (
	UserType UserTypeGroup = "user"
)

type RelationshipTypeGroup = string

const (
	RelationshipType RelationshipTypeGroup = "relationship"
)
