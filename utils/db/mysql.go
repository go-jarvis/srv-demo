package db

import (
	"context"
	"database/sql"

	"github.com/go-jarvis/rum-gonic/rum"
	"github.com/sirupsen/logrus"
)

type DemoDBExecutor string

const (
	Key_DemoDB DemoDBExecutor = "mysql"
)

func WithDBExecutor(db *sql.DB) func(ctx context.Context) context.Context {
	logrus.Infof("注入sqldb")
	return rum.InjectContextValueWith(Key_DemoDB, db)
}

func DBExecutorFrom(ctx context.Context) *sql.DB {
	val := rum.InjectedContextValueFrom(ctx, Key_DemoDB)
	return val.(*sql.DB)
}
