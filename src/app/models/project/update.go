package project

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Update(projectID, ownerID int32, name, url, githubSecret string) error {

	_, err := env.Connect().Exec(`SELECT project.update(
				$1,
				$2,
				$3,
				$4,
				$5
			)
		`,
		projectID,
		name,
		ownerID,
		url,
		githubSecret,
	)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not update project %d, err: %v", projectID, err)
		}

		return err
	}

	return nil
}
