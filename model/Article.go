package model

import (
	"gin_blog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null"`
	Cid     int    `gorm:"type:int;not null"`
	Desc    string `gorm:"type:varchar(200)"`
	Content string `gorm:"type:longtext"`
	Img     string `gorm:"type:varchar(100)"`
}

//添加文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //返回500
	}
	return errmsg.SUCCESS

}

//文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int, int) {
	var artList []Article
	var total int
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&artList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return artList, errmsg.SUCCESS, total
}

//删文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//更新文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

//查询分类下的所有文直
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtList []Article
	var total int
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total

}

//查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}
