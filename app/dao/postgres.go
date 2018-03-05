package dao

import (
	"httpserver-test/config"

	"github.com/go-pg/pg"
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
