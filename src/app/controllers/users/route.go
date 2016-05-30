package users

import (
	"github.com/postgres-ci/app-server/src/app/middleware"
	"github.com/postgres-ci/http200ok"
)

func Route(server *http200ok.Server) {

	server.Get("/users/", listHandler)
	server.Get("/users/get/:UserID/", getHandler)

	server.Post("/users/add/", addHandler)
	server.Post("/users/update/:UserID/", updateHandler)
	server.Post("/users/delete/:UserID/", middleware.CheckPassword, deleteHandler)
}
