package initialize

import (
	"github.com/zzustu/im/global"
	"github.com/zzustu/im/model/data"
	"gopkg.in/ini.v1"
	"path/filepath"
)

var (
	configPath = filepath.Join("resource", "config", "im.ini")
)

// 加载配置文件的配置信息
func loadConfig() error {
	load, err := ini.Load(configPath)
	if err != nil {
		return err
	}

	global.ImCfg = &data.ImCfg{}
	return load.MapTo(global.ImCfg)
}
