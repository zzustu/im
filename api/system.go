package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zzustu/im/global"
	"github.com/zzustu/im/model/view"
	"log"
	"net/http"
)

func Ping(c *gin.Context) {
	_, err := c.Writer.WriteString("PONG")
	if err != nil {
		log.Print(err)
	}
}

func Put(c *gin.Context) {
	global.CM.Broadcast("Hello, 大家好")
}

func Get(c *gin.Context) {
	key := c.Query("key")
	str := global.TB.GetToken(key)
	c.JSON(http.StatusOK, view.NewSuccessResult(str))
}
