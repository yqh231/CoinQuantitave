package log

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	file, err := os.OpenFile("/data/log/monitor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err!=nil{
		panic("open log file, please check /data/log path")
	}

	Logger = log.New(file, "log", log.Ldate|log.Ltime|log.Lshortfile)
}