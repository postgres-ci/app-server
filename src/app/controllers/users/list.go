package users

import (
	"github.com/postgres-ci/app-server/src/app/models/users"
	"github.com/postgres-ci/app-server/src/tools/limit"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/app-server/src/tools/render/pagination"
	"github.com/postgres-ci/http200ok"
)

func listHandler(c *http200ok.Context) {

	var (
		perPage int32 = 20
		query         = c.Request.URL.Query()
	)
	list, err := users.List(perPage, limit.Offset(c, perPage), query.Get("q"))

	if err != nil {

		return
	}

	render.HTML(c, "users/index.html", render.Context{
		"total":      list.Total,
		"users":      list.Users,
		"query":      query.Get("q"),
		"pagination": pagination.New(c, list.Total, perPage),
	})
}
