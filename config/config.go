package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	Log      Log
	Postgres Postgres
	Server   Server
}

type Server struct {
	Port int `default:"8080"`
}

type Postgres struct {
	Ip     string `default:"127.0.0.1"`
	Port   int    `default:"3306"`
	User   string
	Pwd    string
	Dbname string
}

type Log struct {
	Filename string
}

func ParseConfig() (err error) {
	viper.SetConfigName("conf.test")
	viper.AddConfigPath(".")
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("read config file error %v", err)
	}

	if err = viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Printf("config: %v", Conf)

	return err
}
