package password

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/app/models/password"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/params"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func resetHandler(c *http200ok.Context) {

	var (
		currentUser     = c.Get("CurrentUser").(*auth.User)
		newPassword     = c.Request.PostFormValue("new_password")
		confirmPassword = c.Request.PostFormValue("confirm_password")
	)

	if !currentUser.IsSuperuser {

		render.JSONError(c, http.StatusForbidden, "Access denied")

		return
	}

	if newPassword != confirmPassword {

		render.JSONError(c, http.StatusBadRequest, "Entered password not equal confirmed")

		return
	}

	if err := password.Reset(params.ToInt32(c, "UserID"), newPassword); err != nil {

		code := http.StatusInternalServerError

		if e, ok := err.(*errors.Error); ok {

			code = e.Code
		}

		render.JSONError(c, code, err.Error())

		return
	}

	render.JSONok(c)
}
