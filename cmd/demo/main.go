package main

import (
	"github.com/tangx/srv-demo/apis"
	"github.com/tangx/srv-demo/cmd/demo/global"
)

func main() {

	r := global.Server()
	// r.Register(apis.RouterRoot)

	global.Server().Register(apis.RouterRoot)

	global.App.Run(
		r,
	)
}
