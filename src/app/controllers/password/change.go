package password

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/app/models/password"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func changeHandler(c *http200ok.Context) {

	var (
		currentUser     = c.Get("CurrentUser").(*auth.User)
		currentPassword = c.Request.PostFormValue("current_password")
		newPassword     = c.Request.PostFormValue("new_password")
		confirmPassword = c.Request.PostFormValue("confirm_password")
	)

	if newPassword != confirmPassword {

		render.JSONError(c, http.StatusBadRequest, "Your password and confirmation password do not match")

		return
	}

	if err := password.Change(currentUser.ID, currentPassword, newPassword); err != nil {

		code := http.StatusInternalServerError

		if e, ok := err.(*errors.Error); ok {

			code = e.Code
		}

		render.JSONError(c, code, err.Error())

		return
	}

	render.JSONok(c)
}
