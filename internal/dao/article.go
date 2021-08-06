package dao

import (
	"github.com/longjoy/blog-service/internal/model"
	"github.com/longjoy/blog-service/pkg/app"
)

func (d *Dao) CountArticle(title string, state uint8) (int, error){
	article := &model.Article{Title: title, State: state}
	return article.Count(d.engine)
}

func (d *Dao) GetArticle(id uint32) ([]*model.Article, error)  {
	article := &model.Article{
		Model:&model.Model{ID: id},
	}
	return article.Get(d.engine)
}

func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error)  {
	article := &model.Article{Title: title, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateArticle(title, desc, cover_image_url, content, created_by string, state uint8) error {
	article := &model.Article{Title: title,Desc: desc, CoverImageUrl: cover_image_url, Content: content,  State: state, Model:&model.Model{CreatedBy: created_by}}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title, desc, cover_image_url, content, modifiedBy string, state uint8 ) error  {
	article := &model.Article{
		Model:&model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state": state,
		"modified_by": modifiedBy,
		"title": title,
		"desc": desc,
		"cover_image_url": cover_image_url,
		"content": content,
	}
	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error  {
	article := &model.Article{Model:&model.Model{ID: id}}
	return article.Delete(d.engine)
}
