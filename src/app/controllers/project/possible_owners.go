package project

import (
	"github.com/postgres-ci/app-server/src/app/models/project"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func possibleOwnersHendler(c *http200ok.Context) {

	owners, err := project.PossibleOwners()

	if err != nil {

		code := http.StatusInternalServerError

		if e, ok := err.(*errors.Error); ok {

			code = e.Code
		}

		render.JSONError(c, code, err.Error())

		return
	}

	render.JSON(c, owners)
}
