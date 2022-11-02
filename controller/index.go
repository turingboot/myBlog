package controller

import (
	"fmt"
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

	gintemplate.HTML(c, http.StatusOK, "article", gin.H{"posts": posts, "tags": tags})
}

func AdminLoginPage(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "login", nil)
}

func AdminIndexPage(c *gin.Context) {
	fmt.Printf("adminIndexPage\n")
	gintemplate.HTML(c, http.StatusOK, "index", nil)
}
