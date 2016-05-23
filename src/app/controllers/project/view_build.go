package project

import (
	"github.com/postgres-ci/app-server/src/app/models/build"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/params"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"
)

func viewBuildHandler(c *http200ok.Context) {

	build, err := build.View(params.ToInt32(c, "BuildID"))

	if err != nil {

		if errors.IsNotFound(err) {

			render.NotFound(c)

		} else {

			render.InternalServerError(c, err)
		}

		return
	}

	render.HTML(c, "project/view_build.html", render.Context{
		"build": build,
	})
}
