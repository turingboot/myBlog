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
	//posts := make(map[string]string)
	//posts["name"] = "dsadsad"
	gintemplate.HTML(c, http.StatusOK, "article", gin.H{"posts": posts, "tags": tags})
}

func AdminLoginPage(c *gin.Context) {
	//posts := indexPostService.GetPostList()
	//tags := indexTagService.GetTagList()
	//posts := make(map[string]string)
	//posts["name"] = "dsadsad"
	gintemplate.HTML(c, http.StatusOK, "admin_login", nil)
}

func AdminLogin(c *gin.Context) {
	userName := c.PostForm("userName")
	userPwd := c.PostForm("userPwd")
	fmt.Printf("%s\n", userName)
	fmt.Printf("%s", userPwd)
	//gin.H{"flag": true}
	c.JSON(http.StatusOK, gin.H{"flag": true})
	//gintemplate.HTML(c, http.StatusOK, "admin_index", nil)
}

func AdminIndexPage(c *gin.Context) {
	fmt.Printf("adminIndexPage")
	gintemplate.HTML(c, http.StatusOK, "admin_index", nil)
}
