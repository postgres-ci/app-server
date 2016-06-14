package notification

import (
	"github.com/postgres-ci/http200ok"
)

func Route(server *http200ok.Server) {

	server.Get("/notification/method/", getMethodHandler)
	server.Post("/notification/update-method/", updateMethodHandler)
}
