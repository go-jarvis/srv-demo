package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
)

type List struct {
	httpx.MethodGet `path:"/:name/list"`

	Name  string `uri:"name"`
	Money int    `query:"money"`
}

func (l *List) Output(c *gin.Context) (interface{}, error) {

	return l, nil
}
