package app

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/postgres-ci/app-server/src/common"
	"github.com/postgres-ci/app-server/src/env"
	"github.com/postgres-ci/app-server/src/middleware"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func New(config common.Config) *app {

	log.Debugf("Connect to postgresql server. DSN(%s)", config.Connect.DSN())

	connect, err := sqlx.Connect("postgres", config.Connect.DSN())

	if err != nil {

		log.Fatalf("Could not connect to database server: %v", err)
	}

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

func (app *app) route() {

	app.auth()
	app.index()
	app.webhooks()
}

func (app *app) Run() {

	log.Fatal(http.ListenAndServe(app.address, app))
}
