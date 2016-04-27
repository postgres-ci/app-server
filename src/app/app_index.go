package app

import (
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"
)

func (app *app) index() {

	app.Get("/", func(c *http200ok.Context) {

		render.HTML(c, "index.html", nil)
	})
}
