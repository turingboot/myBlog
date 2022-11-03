package models

import "myBlog/common/global"

type Category struct {
	Id    int64  `form:"id" gorm:"primaryKey"`
	Title string `form:"title" gorm:"title"`
}

func AddCateGory(cateGory Category) int64 {
	return global.Db.Table("CateGory").Create(&cateGory).RowsAffected
}

func DelCateGory(id int) Category {
	var cateGory Category
	global.Db.Delete(&cateGory, id)
	return cateGory
}

func UpdateCateGory(cateGory Category) int64 {
	return global.Db.Table("CateGory").Updates(&cateGory).RowsAffected
}

func GetCateGory(id int) Category {
	var cateGory Category
	global.Db.First(&cateGory, id)
	return cateGory
}

func GetCateGoryList() []Category {
	cateGoryList := make([]Category, 0)
	global.Db.Find(&cateGoryList)
	return cateGoryList
}
