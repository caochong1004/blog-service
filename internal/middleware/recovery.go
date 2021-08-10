package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog-service/global"
	"github.com/longjoy/blog-service/pkg/app"
	"github.com/longjoy/blog-service/pkg/email"
	"github.com/longjoy/blog-service/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc  {
	defailtMailer := email.NewEmail(&email.SMTInfo{
		Host: global.EmailSetting.Host,
		Port: global.EmailSetting.Port,
		IsSSL: global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From: global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recovery err : %v"
				global.Logger.WithCallersFrames().Errorf(s, err)
				err := defailtMailer.SendMail(global.EmailSetting.To,
					fmt.Sprintf("异常抛出：发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
					)
				if err != nil {
					global.Logger.Panicf("mail.sendMail err: %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
