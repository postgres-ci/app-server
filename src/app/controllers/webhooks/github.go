package webhooks

import (
	"github.com/postgres-ci/app-server/src/app/models/webhooks/github"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"
)

func githubHandler(c *http200ok.Context) {

	github.Push(github.PushEvent{})

	render.JSON(c, "Webhooks github!")
}
