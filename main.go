package main

import (
	"myBlog/common/initialize"
	"myBlog/common/router"
)

func main() {
	initialize.LoadConfig()
	initialize.MysqlConnector()
	router.Router()

	//var t time.Time
}
