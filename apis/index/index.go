package index

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
	"github.com/tangx/srv-demo/utils/db"
	"golang.org/x/net/context"
)

type Index struct {
	httpx.MethodGet `path:"/index"`
}

func (i *Index) Output(c *gin.Context) (interface{}, error) {

	err := createTable(c)
	if err != nil {
		return nil, fmt.Errorf("create table failed: %v", err)
	}

	return map[string]string{
		"user": "tangx.in",
	}, nil
}

func createTable(ctx context.Context) error {
	db := db.DBExecutorFrom(ctx)

	sql := `
create table if not exists users (
	id int not null primary key auto_increment,
	name varchar(32) not null,
	age int 
) default charset utf8;
`

	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
