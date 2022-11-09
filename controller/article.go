package controller

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleAddPageHandler(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "article_add", nil)
}
func ArticleListHandler(c *gin.Context) {
	//categoryList, errCode := categoryService.GetCategoryList()
	//if errCode != 200 {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status":  errCode,
	//		"message": errmsg.GetErrMsg(errCode),
	//	})
	//	return
	//}
	//
	////这里是为了配合lay-ui的table而配置的特有json模板
	//c.JSON(http.StatusOK, gin.H{
	//	"code":  0,
	//	"msg":   errmsg.GetErrMsg(errCode),
	//	"count": len(categoryList),
	//	"data":  categoryList,
	//})
}
