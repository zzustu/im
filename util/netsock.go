package util

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

// 检测某台主机的某个端口是否可以连通
//
// 参数:
// 		- host: 主机IP地址
// 		- port: 端口号
// 		- timeout: 连接超时时间
// 返回:
//		- bool: 是否可以连通 true:是 false:否
func Reachable(host string, port int, timeout time.Duration) bool {
	if conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", port)), timeout); err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

// 获取本机未被使用的端口
// 从输入的端口开始检测, 如果端口被占用, +1继续检测, 直到找到可用的端口返回
//
// 参数:
// 		- port: 从哪个端口开始检测
// 返回:
//		- int: 可用的端口, 如果有错返回-1
// 		- error: 错误信息
func GetFreePort(port int) (int, error) {
	// 如果输入的端口号大于1<<16, 说明端口号超出范围. (1<<16=65536)
	if port >= 1<<16 {
		return -1, errors.New("port out of range")
	}

	// 如果输入的端口号小于1不报错, 从1<<10端口开始扫描. (1<<10=1024)
	if port < 1 {
		port = 1 << 10
	}

	for Reachable("", port, 1*time.Second) {
		// 如果本机端口可连通, 说明端口号已经被占用, 继续+1看下一个是否已使用
		port = port + 1
		// 如果的端口号大于1<<16, 说明端口号已经遍历到了正常端口号的最大值也没找到. (1<<16=65536)
		if port >= 1<<16 {
			return -1, errors.New("free port not found")
		}
	}
	return port, nil
}

func NICS() {
	itfs, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, itf := range itfs {
		if (itf.Flags & net.FlagUp) == 0 {
			log.Print("网卡未工作")
			continue
		}
		if (itf.Flags & net.FlagLoopback) == 0 {
			log.Print("环回网卡")
			continue
		}
		log.Print(itf)
	}
}
