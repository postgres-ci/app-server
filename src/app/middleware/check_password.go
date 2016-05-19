package middleware

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/app/models/password"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func CheckPassword(c *http200ok.Context) {

	if user, ok := c.Get("CurrentUser").(*auth.User); ok {

		if err := password.Check(user.ID, c.Request.PostFormValue("password")); err == nil {

			return
		}
	}

	render.JSONError(c, http.StatusUnauthorized, "Authentication failed")

	c.Stop()
}
