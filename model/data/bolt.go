package data

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type TokenBolt struct {
	path   string   // 持久化文件存放位置
	bucket []byte   // bucket名称
	db     *bolt.DB // Bolt DB
}

// 新建Token缓存
// 通过key从Token Bucket获得value
//
// 参数:
// 		- path: 持久化文件存放位置
// 		- bucket: bucket名称
// 		- db: Bolt DB
func NewTokenBolt(path string, bucket []byte, db *bolt.DB) *TokenBolt {
	return &TokenBolt{
		path:   path,
		bucket: bucket,
		db:     db,
	}
}

// 通过key从Token Bucket获得value
//
// 参数:
// 		- key: key
func (tb *TokenBolt) GetToken(key string) (value string) {
	_ = tb.db.View(func(tx *bolt.Tx) error {
		get := tx.Bucket(tb.bucket).Get([]byte(key))
		value = fmt.Sprintf("%s", get)
		return nil
	})
	return
}

// 往Token Bucket放入key-value
//
// 参数:
// 		- key: key
// 		- value: value
func (tb *TokenBolt) PutToken(key, value string) error {
	return tb.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(tb.bucket).Put([]byte(key), []byte(value))
	})
}

// 从Token Bucket删除key
//
// 参数:
// 		- key: key
func (tb *TokenBolt) DeleteToken(key string) {
	_ = tb.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(tb.bucket).Delete([]byte(key))
	})
}

// 关闭TokenBolt缓存
func (tb *TokenBolt) Close() error {
	return tb.db.Close()
}
