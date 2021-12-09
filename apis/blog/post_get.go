package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/tangx/srv-demo/utils/db"
)

func init() {
	RouterBlog.Register(&GetPostByID{})
}

type GetPostByID struct {
	httpx.MethodGet `path:"/post/:id"`
	ID              int `uri:"id"`
}

func (po *GetPostByID) Output(c *gin.Context) (interface{}, error) {

	db := db.DBExecutorFrom(c)

	row := db.QueryRow(getPostByIdSql, po.ID)

	post := Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreateTime)
	post.CreateTime = post.CreateTime.Local()
	if err != nil {
		return nil, err
	}

	return post, err
}

var getPostByIdSql = `
select id,title,content,create_time from posts 
  where id = ?;
`
