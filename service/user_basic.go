package service

import (
	"im/helper"
	"im/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码不能为空",
		})
	}
	models.GetUserBasicByAccountPassword(account, helper.GetMd5(password))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
	})
}
