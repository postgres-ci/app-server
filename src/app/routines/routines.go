package routines

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/app/models/build"

	"time"
)

func Run() {

	go func() {

		var (
			authGC  = time.Tick(time.Minute)
			buildGC = time.Tick(10 * time.Minute)
		)

		for {
			select {
			case <-authGC:
				auth.GC()
			case <-buildGC:
				build.GC()
			}
		}
	}()
}
