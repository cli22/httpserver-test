package log

import (
	"log"
	"os"

	"httpserver-test/config"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

var file *os.File

func InitLog(conf config.Config) (err error) {
	// todo append file mode, add ctx
	file, err = os.Create(conf.Log.Filename)
	if err != nil {
		log.Fatalf("open file error %v", err)
	}

	Debug = log.New(file, "Debug:", log.LstdFlags|log.Lshortfile)
	Info = log.New(file, "Info:", log.LstdFlags|log.Lshortfile)
	Warning = log.New(file, "Warning:", log.LstdFlags|log.Lshortfile)
	Error = log.New(file, "Error:", log.LstdFlags|log.Lshortfile)
	return
}
