package apis

import (
	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/go-jarvis/srv-demo/apis/blog"
	"github.com/go-jarvis/srv-demo/apis/index"
	"github.com/go-jarvis/srv-demo/apis/user"
)

var (
	RouterRoot = rum.NewRouterGroup("/")

	RouterV0 = rum.NewRouterGroup("/v0")
)

func init() {

	RouterRoot.Register(&index.Index{})

	RouterRoot.Register(RouterV0)
	{
		RouterV0.Register(user.RouterUser)
		RouterV0.Register(blog.RouterBlog)
	}

}
