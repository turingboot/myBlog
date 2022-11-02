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

func AdminLogin(c *gin.Context) {
	userName := c.PostForm("username")
	userPwd := c.PostForm("password")
	fmt.Printf("%s\n", userName)
	fmt.Printf("%s", userPwd)
	//gin.H{"flag": true}
	c.JSON(http.StatusOK, gin.H{"status": 200})
	//gintemplate.HTML(c, http.StatusOK, "admin_index", nil)
	//c.Redirect(http.StatusOK, "/admin/index")
}

func AdminIndexPage(c *gin.Context) {
	fmt.Printf("adminIndexPage\n")
	gintemplate.HTML(c, http.StatusOK, "index", nil)
}
