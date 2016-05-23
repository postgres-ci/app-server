package users

import (
	"github.com/postgres-ci/app-server/src/app/middleware"
	"github.com/postgres-ci/http200ok"
)

func Route(server *http200ok.Server) {
	server.Get("/users/", listHandler)
	server.Post("/users/add/", addHandler)
	server.Post("/users/edit/:UserID/", updateHandler)
	server.Post("/users/delete/:UserID/", middleware.CheckPassword, deleteHandler)
}
