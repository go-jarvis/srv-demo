package blog

import (
	"database/sql"
	"sync"
	"time"

	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/sirupsen/logrus"
)

var RouterBlog = rum.NewRouterGroup("/blog")

type Post struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
	DeleteTime time.Time `json:"delete_time"`
}

var createTablePosts = `
create table if not exists posts(
	id int not null primary key auto_increment,
	title varchar(256) not null,
	content text,
	create_time timestamp,
	delete_time timestamp null
) default charset=utf8mb4;
`

var once sync.Once

func InitialPost(db *sql.DB) {
	fn := func() {
		logrus.Debugf("try to create table posts")
		_, err := db.Exec(createTablePosts)
		if err != nil {
			logrus.Errorf("create table posts failed: %v", err)
		}
	}

	once.Do(fn)
}
