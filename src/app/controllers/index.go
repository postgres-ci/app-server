package controllers

import (
	"github.com/postgres-ci/app-server/src/app/models/build"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"
)

func Index(server *http200ok.Server) {

	server.Get("/", func(c *http200ok.Context) {

		total, items, err := build.List(1, 0, 20, 0)

		if err == nil {

			render.HTML(c, "index.html", render.Context{
				"total": total,
				"items": items,
			})

			return
		}

		render.HTML(c, "index.html", nil)
	})
}
