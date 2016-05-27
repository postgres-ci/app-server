package project

import (
	"github.com/postgres-ci/app-server/src/app/models/build"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/limit"
	"github.com/postgres-ci/app-server/src/tools/params"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/app-server/src/tools/render/pagination"
	"github.com/postgres-ci/http200ok"
)

func buildsHandler(c *http200ok.Context) {

	var perPage int32 = 20

	list, err := build.List(params.ToInt32(c, "ProjectID"), params.ToInt32(c, "BranchID"), perPage, limit.Offset(c, perPage))

	if err != nil {

		if errors.IsNotFound(err) {

			render.NotFound(c)

		} else {

			render.InternalServerError(c, err)
		}

		return
	}

	render.HTML(c, "project/builds.html", render.Context{
		"branches":    list.Branches,
		"BranchID":    params.ToInt32(c, "BranchID"),
		"ProjectID":   list.ProjectID,
		"ProjectName": list.ProjectName,
		"total":       list.Total,
		"items":       list.Items,
		"pagination":  pagination.New(c, list.Total, perPage),
	})
}
