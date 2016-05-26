package project

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func PossibleOwners() ([]user, error) {

	var owners []user

	err := env.Connect().Select(&owners, `
		SELECT 
			user_id,
			user_name
		FROM project.get_possible_owners()
		`,
	)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not fetch possible owners: %v", err)
		}

		return nil, err
	}

	return owners, nil
}
