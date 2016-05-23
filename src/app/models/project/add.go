package project

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Add(ownerID int32, name, url, githubSecret string) error {

	_, err := env.Connect().Exec(`SELECT project.add(
				$1,
				$2,
				$3,
				$4
			)
		`,
		name,
		ownerID,
		url,
		githubSecret,
	)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not add project: %v", err)
		}

		return err
	}

	return nil
}
