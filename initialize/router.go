package initialize

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zzustu/im/api"
	"github.com/zzustu/im/global"
	"github.com/zzustu/im/util"
	"net/http"
	"time"
)

func getHttpServer() (srv *http.Server, err error) {

	cfg := global.ImCfg.ImServer
	if cfg.HtmlPrefix == "" || cfg.ApiPrefix == "" {
		return nil, errors.New("prefix不能为空")
	}
	if cfg.HtmlPrefix == cfg.ApiPrefix {
		return nil, errors.New("apiprefix与htmlprefix不能相同")
	}

	gin.DefaultWriter = global.LogWriter
	gin.SetMode(cfg.Mode)

	engine := gin.Default()
	engine.Static(cfg.HtmlPrefix, cfg.HtmlPath)

	group := engine.Group(cfg.ApiPrefix)
	addApiRouter(group)

	port, err := util.GetFreePort(cfg.Port)
	if err != nil {
		return nil, err
	}

	srv = &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return srv, nil
}

func addApiRouter(g *gin.RouterGroup) {
	g.GET("/ping", api.Ping)
	g.GET("/get", api.Get)
	g.GET("/put", api.Put)
	g.GET("/im", api.WsConn)

	g.POST("/login", api.Login)
}
