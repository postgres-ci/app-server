package password

import (
	"github.com/postgres-ci/http200ok"
)

func Route(server *http200ok.Server) {

	server.Post("/password/change/", changeHandler)
}
