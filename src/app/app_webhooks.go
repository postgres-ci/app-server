package app

import (
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func (app *app) webhooks() {

	app.Post("/webhooks/native/", func(c *http200ok.Context) {

		token := c.Request.Header.Get("X-Token")

		if len(token) != 36 {

			render.JSONError(c, http.StatusBadRequest, "Invalid X-Token Header")

			return
		}

		render.JSON(c, "Webhooks native!")
	})

	app.Post("/webhooks/github/", func(c *http200ok.Context) {

		render.JSON(c, "Webhooks github!")
	})
}
