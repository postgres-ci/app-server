package webhooks

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/app/models/webhooks/github"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func githubHandler(c *http200ok.Context) {

	defer c.Request.Body.Close()

	log.Debugf("Webhook [GitHub]. Signature: %s, event: %s",
		c.Request.Header.Get("X-Hub-Signature"),
		c.Request.Header.Get("X-GitHub-Event"),
	)

	switch event := c.Request.Header.Get("X-GitHub-Event"); event {
	case "push":

		var push github.PushEvent

		source, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {

			render.JSONError(c, http.StatusInternalServerError, err.Error())

			return
		}

		if err := json.Unmarshal(source, &push); err != nil {

			render.JSONError(c, http.StatusBadRequest, err.Error())

			return
		}

		secret, err := github.Secret(push.Repository.FullName)

		if err != nil {

			if errors.IsNotFound(err) {
				render.JSONError(c, http.StatusNotFound, "Project not nound")
			} else {
				render.JSONError(c, http.StatusInternalServerError, err.Error())
			}

			return
		}

		if secret != "" {

			signature := c.Request.Header.Get("X-Hub-Signature")

			if signature == "" {

				render.JSONError(c, http.StatusForbidden, "Missing X-Hub-Signature header")

				return
			}

			mac := hmac.New(sha1.New, []byte(secret))
			mac.Write(source)

			expectedMAC := hex.EncodeToString(mac.Sum(nil))

			if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {

				render.JSONError(c, http.StatusForbidden, "HMAC verification failed")

				return
			}
		}

		if err := github.Push(push); err != nil {

			render.JSONError(c, http.StatusInternalServerError, err.Error())

			return
		}

	case "ping":

		render.JSON(c, "pong")

		return

	case "":

		render.JSONError(c, http.StatusBadRequest, "Missing X-GitHub-Event header")

		return

	default:

		log.Warnf("GitHub event \"%s\" is not supported", event)
	}

	render.JSONok(c)
}
