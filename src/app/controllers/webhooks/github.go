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

			http.Error(c.Response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		if err := json.Unmarshal(source, &push); err != nil {

			http.Error(c.Response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

			return
		}

		secret, err := github.Secret(push.Repository.FullName)

		if err != nil {

			if errors.IsNotFound(err) {
				http.Error(c.Response, "Project not nound", http.StatusNotFound)
			} else {
				http.Error(c.Response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

			return
		}

		if secret != "" {

			signature := c.Request.Header.Get("X-Hub-Signature")

			if signature == "" {

				http.Error(c.Response, "Missing X-Hub-Signature header", http.StatusForbidden)

				return
			}

			mac := hmac.New(sha1.New, []byte(secret))
			mac.Write(source)

			expectedMAC := hex.EncodeToString(mac.Sum(nil))

			if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {

				http.Error(c.Response, "HMAC verification failed", http.StatusForbidden)

				return
			}
		}

		if err := github.Push(push); err != nil {

			http.Error(c.Response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

	case "":

		http.Error(c.Response, "Missing X-GitHub-Event header", http.StatusBadRequest)

		return

	default:

		log.Warnf("GitHub event \"%s\" is not supported", event)
	}

	render.JSONok(c)
}
