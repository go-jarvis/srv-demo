package main

import (
	"os"

	"github.com/go-jarvis/srv-demo/apis"
	"github.com/go-jarvis/srv-demo/cmd/demo/global"
)

func main() {

	// 数据库初始化
	if os.Getenv("automigrate") == "automigrate" {
		global.Migrate()
	}

	// 启动服务
	r := global.Server()
	r.Register(apis.RouterRoot)
	global.App.Run(
		r,
	)
}
