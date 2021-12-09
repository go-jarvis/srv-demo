package global

import (
	"github.com/go-jarvis/confhttp"
	"github.com/go-jarvis/confmysql"
	"github.com/go-jarvis/jarvis"
	"github.com/tangx/srv-demo/utils/db"
)

var (
	httpserver  = &confhttp.Server{}
	mysqlOnline = &confmysql.Mysql{}
)
var (
	App = jarvis.NewApp()
)

func init() {
	config := &struct {
		HttpServer  *confhttp.Server
		MysqlOnline *confmysql.Mysql
	}{
		HttpServer:  httpserver,
		MysqlOnline: mysqlOnline,
	}

	_ = App.Conf(config)
}

func Server() *confhttp.Server {

	httpserver.WithContextInjectors(
		db.WithDBExecutor(mysqlOnline.DB),
	)

	return httpserver
}
