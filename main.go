package main

import (
	"context"
	"flag"
	"github.com/longjoy/blog-service/global"
	"github.com/longjoy/blog-service/internal/model"
	"github.com/longjoy/blog-service/internal/routers"
	"github.com/longjoy/blog-service/pkg/logger"
	"github.com/longjoy/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	port string
	runMode string
	config string
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

	err  = setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err %v", err)
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
	//信号控制程序优雅的重启
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServer err: %v", err)
		}
	}()
	//等待中断信号
	quit := make(chan os.Signal)
	//接受syscall.SIGINT 和 syscall.sigterm
	signal.Notify(quit,syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shuting down server...")
	//最大时间控制，用于通知该服务端他有5秒的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}
	log.Println("server exiting")

}

func setupSetting() error  {
	setting, err := setting.NewSetting(strings.Split(config,",")...)
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

	//邮件
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil{
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

//启动日志
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

func setupFlag() error  {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode,"mode","", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.Parse()
	return nil
}


