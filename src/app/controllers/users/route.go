package users

import (
	"github.com/postgres-ci/app-server/src/app/middleware"
	"github.com/postgres-ci/http200ok"
)

func Route(server *http200ok.Server) {

	server.Post("/users/delete/:UserID/", middleware.CheckPassword, func(c *http200ok.Context) {

	})
}
