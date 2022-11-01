package service

import (
	"myBlog/common/global"
	"myBlog/models"
)

type TagService struct {
}

func (p *TagService) AddTag(tag models.Tag) int64 {
	return global.Db.Table("tag").Create(&tag).RowsAffected
}

func (p *TagService) DelTag(id int) models.Tag {
	var tag models.Tag
	global.Db.Delete(&tag, id)
	return tag
}

func (p *TagService) UpdateTag(tag models.Tag) int64 {
	return global.Db.Table("tag").Updates(&tag).RowsAffected
}

// GetTag get one tag
func (p *TagService) GetTag(id int) models.Tag {
	var tag models.Tag
	global.Db.First(&tag, id)
	return tag
}

// GetTagList get tag list
func (p *TagService) GetTagList() []models.Tag {
	tagList := make([]models.Tag, 0)
	global.Db.Find(&tagList)
	return tagList
}
