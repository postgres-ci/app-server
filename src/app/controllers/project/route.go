package project

import (
	"github.com/postgres-ci/http200ok"
)

func Route(server *http200ok.Server) {

	server.Get("/project/:ProjectID/builds/", buildsHandler)
}