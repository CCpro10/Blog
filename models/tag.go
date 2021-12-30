package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Model//使用自己定义的model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

//类似于触发器,在新建前设定时间
func (tag *Tag) BeforeCreate( *gorm.DB) error {
	tag.CreatedOn=time.Now().Unix()

	return nil
}

func (tag *Tag) BeforeUpdate( *gorm.DB) error {
	tag.ModifiedOn=time.Now().Unix()

	return nil
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

//传入条件,获取条件筛选后的标签数量
func GetTagTotal(maps interface {}) (count int64){
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) bool{
	db.Create(&Tag {
		Name : name,
		State : state,
		CreatedBy : createdBy,
	})

	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface {}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}