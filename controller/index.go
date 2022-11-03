package controller

import (
	"fmt"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"myBlog/service"
	"net/http"
)

var postService service.PostService
var categoryService service.CategoryService

func Index(c *gin.Context) {
	posts := postService.GetPostList()
	categories, _ := categoryService.GetCategoryList()
	gintemplate.HTML(c, http.StatusOK, "article", gin.H{"posts": posts, "categories": categories})
}

func AdminLoginPage(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "login", nil)
}

func AdminIndexPage(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	gintemplate.HTML(c, http.StatusOK, "index", gin.H{"user_id": id})
}

func AdminUserManagementPage(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "user_manage", nil)
}

func AdminCategoryManagementPage(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "category_manage", nil)
}
