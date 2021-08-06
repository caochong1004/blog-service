package model

import (
	"github.com/jinzhu/gorm"
	"github.com/longjoy/blog-service/pkg/app"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article)TableName() string  {
	return "blog_article"
}


type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}

func (a Article) Count(db *gorm.DB) (int, error)  {
	var count int
	if a.Title != ""{
		db = db.Where("name=?", a.Title)
	}
	err := db.Model(&a).Where("is_del = ? AND state = ?", 0,a.State).Count(&count).Error
	if err != nil{
		return 0, err
	}
	return count, nil
}

func (a Article) Get(db *gorm.DB) ([]*Article, error)  {
	var article []*Article
	err := db.Where("id = ? AND is_del= ?", a.ID, 0).Find(&article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error)  {
	var articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0{
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db.Where("title = ?", a.Title)
	}
	err = db.Where("state = ? AND is_del= ?", a.State, 0).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values interface{}) error  {
	err  := db.Model(&a).Where("is_del = ?", 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error  {
	return db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error
}