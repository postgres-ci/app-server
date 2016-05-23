package app

import (
	"github.com/postgres-ci/app-server/src/app/controllers"
	"github.com/postgres-ci/app-server/src/app/controllers/auth"
	"github.com/postgres-ci/app-server/src/app/controllers/password"
	"github.com/postgres-ci/app-server/src/app/controllers/project"
	"github.com/postgres-ci/app-server/src/app/controllers/users"
	"github.com/postgres-ci/app-server/src/app/controllers/webhooks"
)

func (app *app) route() {

	auth.Route(app.Server)
	users.Route(app.Server)
	password.Route(app.Server)
	project.Route(app.Server)
	webhooks.Route(app.Server)
	controllers.Index(app.Server)
}
