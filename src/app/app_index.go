package app

import (
	"github.com/postgres-ci/app-server/src/app/build"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"
)

func (app *app) index() {

	app.Get("/", func(c *http200ok.Context) {

		total, items, err := build.List(20, 0)

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
