package service

import (
	"myBlog/common/errmsg"
	"myBlog/models"
)

type UserService struct {
}

func (us *UserService) UserinfoInt(ID int) (models.User, int) {
	model := models.GetUserById(ID)
	if model.Username == "" {
		return model, errmsg.ERROR_USER_NOT_EXIST
	}
	return model, errmsg.SUCCSE
}

func (us *UserService) UserinfoStr(userName string) (models.User, int) {
	model := models.GetUserByName(userName)
	if model.Username == "" {
		return model, errmsg.ERROR_USER_NOT_EXIST
	}
	return model, errmsg.SUCCSE
}

func (us *UserService) UserList() ([]models.User, int) {
	modelList := models.GetUserList()
	if modelList == nil {
		return modelList, errmsg.ERROR_USER_NOT_EXIST
	}
	return modelList, errmsg.SUCCSE
}
