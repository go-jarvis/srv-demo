package index

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
)

type Index struct {
	httpx.MethodGet `path:"/index"`
}

func (i *Index) Output(c *gin.Context) (interface{}, error) {

	return map[string]string{
		"user": "tangx.in",
	}, nil
}
