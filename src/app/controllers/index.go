package controllers

import (
	"github.com/postgres-ci/app-server/src/app/models/project"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"
)

func Index(server *http200ok.Server) {

	server.Get("/", func(c *http200ok.Context) {

		items, err := project.List()

		if err != nil {

			render.InternalServerError(c, err)

			return
		}

		render.HTML(c, "index.html", render.Context{
			"items": items,
		})

	})
}
