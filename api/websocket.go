package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zzustu/im/global"
	"github.com/zzustu/im/model/ws"
	"github.com/zzustu/im/util"
	"log"
	"net/http"
)

var (
	upgrade = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func WsConn(c *gin.Context) {

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print(err)
		return
	}

	token := c.Query("token")
	if token == "" {
		_ = conn.WriteJSON(ws.ErrorResponse("请输入Token信息"))
		_ = conn.Close()
	}

	claims, err := util.ParseToken(token)
	if err != nil {
		log.Print(err)
		_ = conn.WriteJSON(ws.ErrorResponse("Token错误"))
		_ = conn.Close()
	}

	if value := global.TB.GetToken(claims.Audience); token != value {
		_ = conn.WriteJSON(ws.ErrorResponse("Token已无效"))
		_ = conn.Close()
	}

	ch := ws.NewChannel(claims.Audience, conn, global.CM)

	global.CM.Connected(ch)

	go ch.Read()
}
