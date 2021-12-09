package apis

import (
	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/tangx/srv-demo/apis/index"
	"github.com/tangx/srv-demo/apis/user"
)

var (
	RouterRoot = rum.NewRouterGroup("/")

	RouterV0 = rum.NewRouterGroup("/v0")
)

func init() {

	RouterRoot.Register(&index.Index{})

	RouterRoot.Register(RouterV0)
	RouterV0.Register(user.RouterUser)
}
