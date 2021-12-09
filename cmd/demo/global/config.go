package global

import (
	"github.com/go-jarvis/confhttp"
	"github.com/go-jarvis/jarvis"
)

var (
	r   = &confhttp.Server{}
	App = jarvis.NewApp()
)

func init() {
	config := &struct {
		HttpServer *confhttp.Server
	}{
		HttpServer: r,
	}
	_ = App.Conf(config)
}

func Server() *confhttp.Server {
	return r
}
