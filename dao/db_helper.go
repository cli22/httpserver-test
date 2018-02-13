package dao

import (
	"github.com/go-pg/pg"
)

var Db = pg.Connect(&pg.Options{
	User:     "postgres",
	Password: "",
	Database: "test",
})

func init() {
	err := createSchema(Db)
	if err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&srv.User{}, &srv.Relationship{}} {
		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

type User struct {
	Id   int64
	Uid  int64
	Name string
	Type string
}

type Relationship struct {
	Id         int64
	Uid        int64
	AnotherUid int64
	State      RelationshipState
	Type       RelationshipType
}

type RelationshipState string

const (
	Liked    RelationshipState = "liked"
	Disliked RelationshipState = "disliked"
	Matched  RelationshipState = "matched"
	Default  RelationshipState = ""
)

type RelationshipType string

const (
	relationship RelationshipType = "relationship"
)
