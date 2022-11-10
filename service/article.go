package service

import (
	"myBlog/common/errmsg"
	"myBlog/common/global"
	"myBlog/dal/model"
	"myBlog/dal/query"
)

type ArticleService struct {
}

func (as *ArticleService) Add(ar model.Article, ard model.ArticleDetail) int {
	var a = query.Use(global.Db).Article
	var ad = query.Use(global.Db).ArticleDetail
	if a.Create(&ar) != nil {
		return errmsg.ERROR
	}
	if ad.Create(&ard) != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func (as *ArticleService) GetList() ([]*model.Article, int) {
	var a = query.Use(global.Db).Article
	aList, err := a.Find()
	if err != nil {
		return aList, errmsg.ERROR
	}
	return aList, errmsg.ERROR
}
