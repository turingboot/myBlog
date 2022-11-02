package controller

import (
	"github.com/gin-gonic/gin"
	"myBlog/common/errmsg"
	"myBlog/service"
	"net/http"
	"strconv"
)

var userService service.UserService

func UserDetailModifyHandler(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	user, errCode := userService.UserinfoInt(ID)
	if errCode != 200 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"Status": errCode,
			"Msg":    errmsg.GetErrMsg(errCode),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": errCode,
		"data":   user,
		"Msg":    errmsg.GetErrMsg(errCode),
	})
}

func UserLoginHandler(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	user, errCode := userService.UserinfoStr(username)
	if errCode != 200 {
		c.JSON(http.StatusOK, gin.H{
			"status":  errCode,
			"message": errmsg.GetErrMsg(errCode),
		})
		return
	}

	if user.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_PASSWORD_WRONG,
			"message": errmsg.GetErrMsg(errmsg.ERROR_PASSWORD_WRONG),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  errCode,
		"data":    user,
		"message": errmsg.GetErrMsg(errCode),
	})
}
