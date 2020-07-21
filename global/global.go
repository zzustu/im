package global

import (
	"github.com/zzustu/im/model/data"
	"github.com/zzustu/im/model/ws"
	"io"
	"xorm.io/xorm"
)

var (
	Version   string          // 版本
	DB        *xorm.Engine    // 数据库
	ImCfg     *data.ImCfg     // 全局配置信息
	TB        *data.TokenBolt // BoltDB
	LogWriter io.Writer       // 日志文件
	CM        = ws.NewChannelManager()
)
