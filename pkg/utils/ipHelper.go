package utils

import (
	_ "fmt"
	"strings"
	"sync"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var (
	searcher *xdb.Searcher
	once     sync.Once
)

// InitIPDB 初始化 IP 库
func InitIPDB(dbPath string) {
	once.Do(func() {
		var err error
		cBuff, err := xdb.LoadContentFromFile(dbPath)
		if err != nil {
			RecordError("加载 xdb 内容失败：", err)
		}

		searcher, err = xdb.NewWithBuffer(xdb.IPv4, cBuff)
		if err != nil {
			RecordError("初始化ip库失败：", err)
		}
	})
}

// GetLocation 通过 IP 获取位置字符串
// 返回格式示例: "中国 上海" 或 "美国"
func GetLocation(ip string) string {
	if searcher == nil {
		return "未知位置(DB未加载)"
	}

	// 本地开发环境IP
	if ip == "127.0.0.1" || ip == "::1" {
		return "本地局域网"
	}

	// 查询
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		return "未知位置"
	}

	// ip2region 返回格式: 国家|区域|省份|城市|ISP
	// 中国|0|上海|上海市|电信
	parts := strings.Split(region, "|")
	var location string

	// 去0
	if parts[0] != "0" {
		location += parts[0] + " " // 国家
	}
	if parts[2] != "0" {
		location += parts[2] + " " // 省份
	}
	if parts[3] != "0" && parts[3] != parts[2] {
		location += parts[3] // 城市
	}

	// 去除首尾空格
	return strings.TrimSpace(location)
}
