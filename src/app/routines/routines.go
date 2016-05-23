package routines

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/env"

	"time"
)

func Run() {

	go func() {

		gc := time.Tick(10 * time.Minute)

		for {

			if _, err := env.Connect().Exec("SELECT auth.gc()"); err != nil {

				log.Errorf("Auth.GC %v", err)
			}

			<-gc
		}
	}()
}
