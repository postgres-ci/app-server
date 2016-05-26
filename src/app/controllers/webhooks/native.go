package webhooks

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/app/models/webhooks"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/hooks/git"
	"github.com/postgres-ci/http200ok"

	"encoding/json"
	"net/http"
)

func nativeHandler(c *http200ok.Context) {

	defer c.Request.Body.Close()

	log.Debugf("Webhook [native]. Token: %s, event: %s",
		c.Request.Header.Get("X-Token"),
		c.Request.Header.Get("X-Event"),
	)

	token := c.Request.Header.Get("X-Token")

	if len(token) != 36 {

		render.JSONError(c, http.StatusBadRequest, "Invalid X-Token Header")

		return
	}

	switch event := c.Request.Header.Get("X-Event"); event {

	case "commit":

		var commit struct {
			Ref string `json:"ref"`
			git.Commit
		}

		if err := json.NewDecoder(c.Request.Body).Decode(&commit); err != nil {

			render.JSONError(c, http.StatusBadRequest, "Json error: %v", err)

			return
		}

		if err := webhooks.Commit(c.Request.Header.Get("X-Token"), commit.Ref, commit.Commit); err != nil {

			render.JSONError(c, http.StatusBadRequest, "Commit error: %v", err)

			return
		}

		render.JSONok(c)

	case "push":

		var push git.Push

		if err := json.NewDecoder(c.Request.Body).Decode(&push); err != nil {

			render.JSONError(c, http.StatusBadRequest, "Json error: %v", err)

			return
		}

		if err := webhooks.Push(c.Request.Header.Get("X-Token"), push); err != nil {

			render.JSONError(c, http.StatusBadRequest, "Push error: %v", err)

			return
		}

		render.JSONok(c)

	default:

		render.JSONError(c, http.StatusBadRequest, "Unreachable X-Event: %s", event)
	}
}
