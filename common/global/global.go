package global

import (
	"gorm.io/gorm"
	"myBlog/common/config"
)

var (
	Config config.Config
	Db     *gorm.DB
)
