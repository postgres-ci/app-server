package project

import (
	"github.com/postgres-ci/app-server/src/app/middleware"
	"github.com/postgres-ci/http200ok"
)

func Route(server *http200ok.Server) {

	server.Get("/project/:ProjectID/builds/", buildsHandler)
	server.Get("/project/:ProjectID/builds/branch/:BranchID/", buildsHandler)
	server.Get("/project/:ProjectID/build/:BuildID/", viewBuildHandler)
	server.Post("/project/add/", addHandler)
	server.Get("/project/:ProjectID/get/", getHandler)
	server.Post("/project/update/:ProjectID/", updateHandler)
	server.Post("/project/delete/:ProjectID/", middleware.CheckPassword, deleteHandler)
}
