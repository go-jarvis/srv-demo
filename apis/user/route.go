package user

import (
	"github.com/go-jarvis/rum-gonic/rum"
)

var (
	RouterUser = rum.NewRouterGroup("/user")
)

func init() {
	RouterUser.Register(&List{})
}
