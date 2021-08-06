package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/longjoy/blog-service/pkg/app"
)

type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8
}

func (t Tag)TableName() string  {
	return "blog_tag"
}

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}

func (t Tag) Count(db *gorm.DB) (int, error)  {
	var count int
	if t.Name != ""{
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?",t.State)
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil{
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error)  {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0{
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil{
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error  {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error  {
	err  := db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error  {
	return db.Where("id = ? AND is_del = ?", t.ID, 0).Delete(&t).Error
}

func (t Tag) Get(db *gorm.DB) ([]*Tag, error)  {
	var tags []*Tag
	err := db.Where("id = ? AND is_del= ?", t.ID, 0).Find(&tags).Error
	if err != nil {
		fmt.Println("err=", err)
		return nil, err
	}
	return tags, nil
}

