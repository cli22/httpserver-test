package dao

import (
	"github.com/go-pg/pg"

	"httpserver-test/config"
)

var Db *pg.DB
// add Timeout
func NewPg(conf config.Config) *pg.DB {
	var db *pg.DB
	db = pg.Connect(&pg.Options{
		User:     conf.Postgres.User,
		Password: conf.Postgres.Pwd,
		Database: conf.Postgres.Dbname,
	})

	return db

}
