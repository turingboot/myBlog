package service

import (
	"myBlog/common/errmsg"
	"myBlog/models"
)

type CategoryService struct {
}

func (cs *CategoryService) GetCategoryList() ([]models.Category, int) {
	cgList := models.GetCateGoryList()
	if len(cgList) < 1 {
		return cgList, errmsg.ERROR_CATE_NOT_EXIST
	}
	return cgList, errmsg.SUCCSE
}
