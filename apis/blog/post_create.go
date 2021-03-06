package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/go-jarvis/srv-demo/utils/db"
)

func init() {
	RouterBlog.Register(&CreatePost{})
}

type CreatePost struct {
	httpx.MethodPost `path:"/post"`

	Body Post `body:"body"`
}

func (po *CreatePost) Output(c *gin.Context) (interface{}, error) {

	post := po.Body

	db := db.DBExecutorFrom(c)

	_, err := db.Exec(insertPostSql, post.Title, post.Content)
	if err != nil {
		return nil, err
	}
	return "success", nil
}

var insertPostSql = `
insert into posts(title, content, create_time)
	values(?, ?, Now());
`
