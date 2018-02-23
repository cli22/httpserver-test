package log

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.Create("httpserver.log")
	if err != nil {
		log.Fatalln("open file error!")
	}
	Info = log.New(file, "Info:", log.LstdFlags|log.Lshortfile)
	Warning = log.New(file, "Warning:", log.LstdFlags|log.Lshortfile)
	Error = log.New(file, "Error·:", log.LstdFlags|log.Lshortfile)

}
