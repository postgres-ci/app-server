package users

import (
	"github.com/postgres-ci/app-server/src/app/models/users"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/params"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func getHandler(c *http200ok.Context) {

	user, err := users.Get(params.ToInt32(c, "UserID"))

	if err != nil {

		code := http.StatusInternalServerError

		if e, ok := err.(*errors.Error); ok {

			code = e.Code
		}

		render.JSONError(c, code, err.Error())

		return
	}

	render.JSON(c, user)
}
