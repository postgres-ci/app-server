package users

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/app/models/users"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
	"strings"
)

func addHandler(c *http200ok.Context) {

	currentUser := c.Get("CurrentUser").(*auth.User)

	if !currentUser.IsSuperuser {

		render.JSONError(c, http.StatusForbidden, "Access denied")

		return
	}

	var (
		login       = strings.TrimSpace(c.Request.PostFormValue("login"))
		password    = strings.TrimSpace(c.Request.PostFormValue("password"))
		name        = strings.TrimSpace(c.Request.PostFormValue("name"))
		email       = strings.TrimSpace(c.Request.PostFormValue("email"))
		isSuperuser bool
	)

	if login == "" || password == "" || email == "" {

		render.JSONError(c, http.StatusOK, "\"Login\", \"password\" and \"email\" fields are required")

		return
	}

	if on := c.Request.PostFormValue("is_superuser"); on != "" {

		isSuperuser = true
	}

	if err := users.Add(login, password, name, email, isSuperuser); err != nil {

		code := http.StatusInternalServerError

		if e, ok := err.(*errors.Error); ok {

			code = e.Code
		}

		render.JSONError(c, code, err.Error())

		return
	}

	render.JSONok(c)
}
