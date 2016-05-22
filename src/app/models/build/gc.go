package build

import (
	"github.com/postgres-ci/app-server/src/env"
)

func GC() {

	env.Connect().Exec("SELECT build.gc()")
}
