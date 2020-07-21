package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zzustu/im/global"
	"github.com/zzustu/im/util"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "请输入用户名和密码",
		})
		return
	}

	token, err := util.NewToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": "生成Token失败",
		})
		return
	}

	err = global.TB.PutToken(username, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": "存放Token错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}
