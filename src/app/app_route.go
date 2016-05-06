package app

import (
	"github.com/postgres-ci/app-server/src/app/controllers"
	"github.com/postgres-ci/app-server/src/app/controllers/auth"
	"github.com/postgres-ci/app-server/src/app/controllers/webhooks"
)

func (app *app) route() {

	auth.Route(app.Server)
	webhooks.Route(app.Server)
	controllers.Index(app.Server)
}
