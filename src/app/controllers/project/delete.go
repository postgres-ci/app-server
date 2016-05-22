package project

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/app/models/project"
	"github.com/postgres-ci/app-server/src/tools/params"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func deleteHandler(c *http200ok.Context) {

	currentUser := c.Get("CurrentUser").(*auth.User)

	if !currentUser.IsSuperuser {

		render.JSONError(c, http.StatusForbidden, "Access denied")

		return
	}

	if err := project.Delete(params.ToInt32(c, "ProjectID")); err != nil {

		render.JSONError(c, http.StatusInternalServerError, err.Error())

		return
	}

	render.JSONok(c)
}
