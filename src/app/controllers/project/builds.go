package project

import (
	"github.com/postgres-ci/app-server/src/app/models/build"
	"github.com/postgres-ci/app-server/src/tools/params"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"
)

func buildsHandler(c *http200ok.Context) {

	list, err := build.List(params.ToInt32(c, "ProjectID"), params.ToInt32(c, "BranchID"), 10, 0)

	if err != nil {

		return
	}

	render.HTML(c, "project/builds.html", render.Context{
		"branches": list.Branches,
		"total":    list.Total,
		"items":    list.Items,
	})
}
