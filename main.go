package main

import (
	"github.com/longjoy/blog-service/global"
	"github.com/longjoy/blog-service/internal/model"
	"github.com/longjoy/blog-service/internal/routers"
	"github.com/longjoy/blog-service/pkg/logger"
	"github.com/longjoy/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init()  {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
	err = setupDBEngine()
	if err != nil{
		log.Fatalf("init.setupDBEngine err %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用go 做项目
// @termsOfService https://github.com/longjoy/blog-service/
func main()  {
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadHeaderTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}

	s.ListenAndServe()
}

func setupSetting() error  {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil{
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil{
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	//jwt
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil{
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error  {
	fileName := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: fileName,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	},"", log.LstdFlags).WithCaller(2)
	return nil
}

func setupDBEngine() error  {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
