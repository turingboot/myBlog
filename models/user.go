package models

import (
	"gorm.io/gorm"
	"myBlog/common/global"
)

//type User struct {
//	gorm.Model
//	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
//	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
//	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
//}

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username"  label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password"  label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role"  label:"角色码"`
}

func GetUserById(ID int) User {
	var user User
	global.Db.Where("id = ?", ID).First(&user)
	return user
}

func GetUserByName(username string) User {
	var user User
	global.Db.Where("username = ?", username).First(&user)
	return user
}
