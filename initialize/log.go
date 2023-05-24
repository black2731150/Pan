package initialize

import (
	"log"
	"pan/global"
	"pan/pkg/logger"

	"gopkg.in/natefinch/lumberjack.v2"
)

func initLogeer() error {
	path := "logs/"
	name := "panlog"
	ext := ".log"
	filename := path + name + ext

	global.Panserver.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  filename,
		MaxSize:   300,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func InitLog() {
	err := initLogeer()
	if err != nil {
		log.Fatalf("initLogeer err:%v", err)
	}
}
