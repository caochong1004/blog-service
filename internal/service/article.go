package service

import (
	"github.com/longjoy/blog-service/internal/model"
	"github.com/longjoy/blog-service/pkg/app"
)

type CountArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type GetArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
	Desc string `form:"desc" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=3,max=100"`
	Content string `form:"content" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
	Title string `form:"title" binding:"required,min=3,max=100"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
	Desc string `form:"desc" binding:"required,min=3,max=100"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=3,max=100"`
	Content string `form:"content" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}


func (svc *Service) CountArticle(param *CountArticleRequest) (int, error)  {
	return svc.dao.CountArticle(param.Title, param.State)
}

func (svc *Service) GetArticle(param *GetArticleRequest) ([]*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}

func (svc *Service) ArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error)  {
	return svc.dao.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error  {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.CoverImageUrl, param.Content, param.CreatedBy, param.State)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error  {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.CoverImageUrl, param.Content,param.ModifiedBy, param.State)
}

func (svc *Service) DeleteArticle(param *DeleteRequest) error  {
	return svc.dao.DeleteArticle(param.ID)
}


