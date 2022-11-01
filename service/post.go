package service

import (
	"myBlog/common/global"
	"myBlog/models"
)

type PostService struct {
}

// AddPost 添加
func (p *PostService) AddPost(post models.Post) int64 {
	return global.Db.Table("post").Create(&post).RowsAffected
}

// DelPost 删除
func (p *PostService) DelPost(id int) int64 {
	return global.Db.Delete(&models.Post{}, id).RowsAffected
}

// UpdatePost 修改
func (p *PostService) UpdatePost(post models.Post) int64 {
	return global.Db.Updates(&post).RowsAffected
}

// GetPost get
func (p *PostService) GetPost(id int) models.Post {
	var post models.Post
	global.Db.First(&post, id)
	return post
}

// GetPostList get post list
func (p *PostService) GetPostList() []models.Post {
	postList := make([]models.Post, 0)
	global.Db.Find(&postList)
	return postList
}
