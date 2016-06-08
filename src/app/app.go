package app

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/postgres-ci/app-server/src/app/middleware"
	"github.com/postgres-ci/app-server/src/app/routines"
	"github.com/postgres-ci/app-server/src/common"
	"github.com/postgres-ci/app-server/src/env"
	"github.com/postgres-ci/http200ok"

	"net/http"
	"os"
	"runtime"
)

func New(config common.Config) *app {

	connect, err := sqlx.Connect("postgres", config.Connect.DSN())

	if err != nil {

		log.Fatalf("Could not connect to database server: %v", err)
	}

	log.Debugf("Connect to postgresql server. DSN(%s)", config.Connect.DSN())

	env.SetConnect(connect)

	app := &app{
		address: config.Address,
		Server:  http200ok.New(),
	}

	app.Use(middleware.Authentication)
	app.route()

	return app
}

type app struct {
	*http200ok.Server
	address string
}

func (app *app) Run() {

	routines.Run()

	log.Infof("Postgres-CI app-server running on address: %s, pid: %d", app.address, os.Getpid())
	log.Debugf("Runtime version: %s", runtime.Version())

	log.Fatal(http.ListenAndServe(app.address, app.Server))
}
