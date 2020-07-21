package initialize

import (
	"github.com/boltdb/bolt"
	"github.com/zzustu/im/global"
	"github.com/zzustu/im/model/data"
	"os"
	"time"
)

// 初始化BoltDB
func loadBoltDB() error {
	cfg := global.ImCfg.ImBolt
	db, err := bolt.Open(cfg.Path, os.ModePerm, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		return err
	}

	bucket := []byte(cfg.Bucket)
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucket)
		return err
	})
	if err != nil {
		return err
	}

	global.TB = data.NewTokenBolt(cfg.Path, bucket, db)

	return nil
}
