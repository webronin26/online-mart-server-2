package main

import (
	"online-mart-server-2/config"
	"online-mart-server-2/pkg/server"
	"online-mart-server-2/pkg/store"
)

func main() {
	// initial server config
	// 初始化設定檔案
	config.Init()

	// initial server database
	// 初始化資料庫
	store.Init()

	// initial server middleware & router
	// 初始化伺服器 : 中間鍵 ＆ 路由
	server.Init()
}
