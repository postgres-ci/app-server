package project

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/app/models/project"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
	"strconv"
	"strings"
)

func addHandler(c *http200ok.Context) {

	currentUser := c.Get("CurrentUser").(*auth.User)

	if !currentUser.IsSuperuser {

		render.JSONError(c, http.StatusForbidden, "Access denied")

		return
	}

	var (
		name         = strings.TrimSpace(c.Request.PostFormValue("name"))
		url          = strings.TrimSpace(c.Request.PostFormValue("url"))
		githubSecret = strings.TrimSpace(c.Request.PostFormValue("github_secret"))
		ownerID      int32
	)

	if value, err := strconv.ParseInt(c.Request.PostFormValue("owner_id"), 10, 32); err == nil {

		ownerID = int32(value)
	}

	if err := project.Add(ownerID, name, url, githubSecret); err != nil {

		code := http.StatusInternalServerError

		if e, ok := err.(*errors.Error); ok {

			code = e.Code
		}

		render.JSONError(c, code, err.Error())

		return
	}

	render.JSONok(c)
}
