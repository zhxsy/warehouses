package app

import (
	"errors"
	"github.com/cfx/warehouses/helper"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var log *logrus.Logger

func InitLog(apEnv, pathDir string) {
	log = logrus.New()
	path := "./logs"
	if envPath, ok := GetEnvPath(apEnv); ok {
		path = envPath + pathDir
	}

	filename := path + "/" + ProjectName + ".log"

	_, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		cp, _ := helper.GetCurrentPath()
		if _, err = os.Stat(cp + path); err != nil && os.IsNotExist(err) {
			log.Println("doing mkdir...: " + cp + path)
			if err = os.MkdirAll(cp+path, os.ModePerm); err != nil {
				panic(errors.New("init_log_path err: " + err.Error()))
			}
		}
		filename = cp + filename
		_, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			panic(errors.New("init_log_err err:" + err.Error()))
		}
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    0x200, // 单个文件最大 10M
		MaxBackups: 30,    // 文件最大保存 30 个
		MaxAge:     30,    // 文件最大保存 30 天
		LocalTime:  true,
		Compress:   true,
	}

	log.SetOutput(lumberjackLogger)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: DatetimeFmt,
		DataKey:         "fields",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "level",
		},
	})

	if IsPrd() {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.TraceLevel)
	}
	log.Hooks.Add(NewContextHook())
	log.WithField("path", path).Info("log_init_success")
}

func Log() *logrus.Logger {
	return log
}
