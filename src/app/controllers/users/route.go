package users

import (
	"github.com/postgres-ci/app-server/src/app/middleware"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/app-server/src/tools/render/pagination"
	"github.com/postgres-ci/http200ok"
)

func Route(server *http200ok.Server) {

	server.Get("/users/", func(c *http200ok.Context) {

		render.HTML(c, "users/index.html", render.Context{
			"pagination": pagination.New(c, 2562, 20),
		})
	})

	server.Post("/users/delete/:UserID/", middleware.CheckPassword, func(c *http200ok.Context) {

	})
}
