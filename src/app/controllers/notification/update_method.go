package notification

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/app/models/notification"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func updateMethodHandler(c *http200ok.Context) {

	var (
		currentUser = c.Get("CurrentUser").(*auth.User)
		method      = c.Request.PostFormValue("method")
		textID      = c.Request.PostFormValue("text_id")
	)

	if err := notification.UpdateMethod(currentUser.ID, method, textID); err != nil {

		code := http.StatusInternalServerError

		if e, ok := err.(*errors.Error); ok {

			code = e.Code
		}

		render.JSONError(c, code, err.Error())

		return
	}

	render.JSONok(c)

}
