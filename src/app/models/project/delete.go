package project

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/env"
)

func Delete(projectID int32) error {

	if _, err := env.Connect().Exec(`SELECT project.delete($1)`, projectID); err != nil {

		log.Errorf("Could not delete project %d, err: %v", projectID, err)

		return err
	}

	return nil
}
