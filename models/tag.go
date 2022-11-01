package models

type Tag struct {
	Id    int64  `form:"id" gorm:"primaryKey"`
	Title string `form:"title" gorm:"title"`
}
