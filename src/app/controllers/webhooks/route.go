package webhooks

import (
	"github.com/postgres-ci/http200ok"
)

func Route(server *http200ok.Server) {

	server.Post("/webhooks/native/", nativeHandler)
	server.Post("/webhooks/github/", githubHandler)
}
