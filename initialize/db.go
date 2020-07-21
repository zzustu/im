package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zzustu/im/global"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

// 连接数据库
func connectDB() (err error) {

	// 读取配置信息
	db := global.ImCfg.ImDB

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.Username, db.Password, db.Host, db.Port, db.Schema)

	global.DB, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return
	}

	global.DB.ShowSQL(db.ShowSQL)

	global.DB.SetLogger(log.NewSimpleLogger(global.LogWriter))

	return global.DB.Ping()
}
