package controller

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"myBlog/common/errmsg"
	"myBlog/dal/model"
	"myBlog/service"
	"net/http"
)

type ArticleInput struct {
	ArticleID      string
	ArticleUser    string `form:"article_user"`
	Title          string `form:"title"`
	TopFlag        string `form:"top_flag"`
	ArticleSummary string `form:"article_summary"`
	ContentMd      string `form:"content_md"`
	ContentHTML    string `form:"content_html"`
}

var articleService service.ArticleService

func ArticleAddPageHandler(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "article_add", nil)
}

func ArticleAddHandler(c *gin.Context) {
	var inputA ArticleInput

	if c.ShouldBind(&inputA) != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "数据绑定错误",
		})
		return
	}

	articleID := uuid.New().String()
	ar := model.Article{
		ArticleID:      articleID,
		ArticleUser:    inputA.ArticleUser,
		Title:          inputA.Title,
		TopFlag:        inputA.TopFlag,
		ArticleSummary: inputA.ArticleSummary,
	}
	arD := model.ArticleDetail{
		ArticleID:   articleID,
		ContentMd:   inputA.ContentMd,
		ContentHTML: inputA.ContentHTML,
	}

	errCode := articleService.Add(ar, arD)
	if errCode != 200 {
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"message": errmsg.GetErrMsg(errCode),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"message": errmsg.GetErrMsg(errCode),
	})
}

func ArticleListHandler(c *gin.Context) {
}
