package controller

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"myBlog/service"
	"net/http"
)

var indexPostService service.PostService
var indexTagService service.TagService

func Index(c *gin.Context) {
	posts := indexPostService.GetPostList()
	tags := indexTagService.GetTagList()
	//posts := make(map[string]string)
	//posts["name"] = "dsadsad"
	gintemplate.HTML(c, http.StatusOK, "article", gin.H{"posts": posts, "tags": tags})
}

func AdminIndex(c *gin.Context) {
	//posts := indexPostService.GetPostList()
	//tags := indexTagService.GetTagList()
	//posts := make(map[string]string)
	//posts["name"] = "dsadsad"
	gintemplate.HTML(c, http.StatusOK, "admin_login", nil)
}
