package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/longjoy/blog-service/internal/service"
	"github.com/longjoy/blog-service/pkg/app"
	"github.com/longjoy/blog-service/pkg/convert"
	"github.com/longjoy/blog-service/pkg/errcode"
)

type Article struct {}

func NewArticle() Article  {
	return Article{}
}

func (a Article) Get(c *gin.Context)  {
	param := service.GetArticleRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}
	response.ToResponse(article)
	return
}

func (a Article) List(c *gin.Context)  {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountArticle(&service.CountArticleRequest{Title: param.Title,
		State: param.State})

	if err != nil {
		//global.Logger.Errorf("svc.CountTag err: %v ", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}
	articles, err  := svc.ArticleList(&param, &pager)
	if err != nil{
		//global.Logger.Errorf("svc.GetTagList err: %v ", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(articles, totalRows)
	return
}

func (a Article) Create(c *gin.Context)  {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (a Article) Update(c *gin.Context)  {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (a Article) Delete(c *gin.Context)  {
	param := service.DeleteRequest{ID: convert.StrTo(c.Param("id")).MustUint32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
