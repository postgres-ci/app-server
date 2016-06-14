package notification

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/app/models/notification"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func getMethodHandler(c *http200ok.Context) {

	method, err := notification.GetMethod(c.Get("CurrentUser").(*auth.User).ID)

	if err != nil {

		code := http.StatusInternalServerError

		if e, ok := err.(*errors.Error); ok {

			code = e.Code
		}

		render.JSONError(c, code, err.Error())

		return
	}

	render.JSON(c, method)
}
