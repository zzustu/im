package initialize

import (
	"github.com/zzustu/im/global"
	"io"
	"log"
	"os"
)

func configLog() error {
	cfg := global.ImCfg.ImLog
	f, err := os.OpenFile(cfg.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	global.LogWriter = io.MultiWriter(os.Stdout, f)
	log.SetPrefix(cfg.Prefix + " ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(global.LogWriter)
	return nil
}
