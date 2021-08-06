package api

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog-service/internal/service"
	"github.com/longjoy/blog-service/pkg/app"
	"github.com/longjoy/blog-service/pkg/errcode"
)

func GetAuth(c *gin.Context)  {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil{
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil{
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"token":token})
}
