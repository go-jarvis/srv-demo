package db

import (
	"context"
	"database/sql"

	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/sirupsen/logrus"
)

type DemoDBExecutor string

const (
	Key_DemoOnlineDB DemoDBExecutor = "mysql-online"
)

func WithOnlineDBExecutor(db *sql.DB) func(ctx context.Context) context.Context {
	logrus.Infof("注入sqldb")
	return rum.WithContextValue(Key_DemoOnlineDB, db)
}

func OnlineDBExecutorFrom(ctx context.Context) *sql.DB {
	val := rum.FromContextValue(ctx, Key_DemoOnlineDB)
	return val.(*sql.DB)
}
