package router

import (
	"fmt"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"html/template"
	"myBlog/common/global"
	"myBlog/controller"
	"myBlog/templates/functions"
	"net/http"
)

func Router() {
	engine := gin.Default()

	// 静态资源请求映射

	engine.Static("/assets", "./assets")
	engine.StaticFS("/static", http.Dir("./static"))
	functions := template.FuncMap{
		"pYear":  functions.PYear,
		"pMonth": functions.PMonth,
		"pDay":   functions.PDay,
	}
	engine.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "templates/frontend",
		Extension:    ".html",
		Master:       "article",
		Funcs:        functions,
		DisableCache: true,
	})

	// 前台
	engine.GET("/", controller.Index)
	engine.GET("/index", controller.Index)

	mw := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
		Root:      "templates/backend",
		Extension: ".html",
		//Master:       "login",
		DisableCache: true,
	})

	// 后台管理员前端接口
	web := engine.Group("/admin", mw)

	{
		web.GET("/", controller.AdminLoginPage)
		web.GET("/login", controller.AdminLoginPage)
		web.GET("/index/:id", controller.AdminIndexPage)
		web.GET("/user", controller.AdminUserManagementPage)
		web.GET("/category", controller.AdminCategoryManagementPage)

	}

	{
		//
		web.POST("/doLogin", controller.UserLoginHandler)
		//web.GET("/user/detail/:username", controller.AdminIndexPage)
		web.GET("/user/all", controller.UserListHandler)
	}
	{
		//
		web.GET("/category/all", controller.CategoryListHandler)
	}

	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}
