package service

import (
	"myBlog/common/errmsg"
	"myBlog/common/global"
	"myBlog/dal/model"
	"myBlog/dal/query"
)

type UserService struct {
}

func (us *UserService) UserinfoInt(ID int) (model.User, int) {
	var u = query.Use(global.Db).User
	user, _ := u.GetUserById(ID)
	if user.Username == "" {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCSE
}

func (us *UserService) UserinfoStr(userName string) (model.User, int) {
	var u = query.Use(global.Db).User
	user, _ := u.GetUserByName(userName)
	if user.Username == "" {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCSE
}

func (us *UserService) UserList() ([]model.User, int) {
	var u = query.Use(global.Db).User
	userList, _ := u.GetUserList()
	if userList == nil {
		return userList, errmsg.ERROR_USER_NOT_EXIST
	}
	return userList, errmsg.SUCCSE
}
