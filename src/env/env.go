package env

import (
	"github.com/jmoiron/sqlx"

	"time"
)

var connect *sqlx.DB

const (
	MaxOpenConns    = 5
	MaxIdleConns    = 2
	ConnMaxLifetime = time.Minute * 30
)

func SetConnect(c *sqlx.DB) {

	connect = c

	connect.SetMaxOpenConns(MaxOpenConns)
	connect.SetMaxIdleConns(MaxIdleConns)
	//	connect.SetConnMaxLifetime(ConnMaxLifetime)
}

func Connect() *sqlx.DB {

	return connect
}
