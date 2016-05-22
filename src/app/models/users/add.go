package users

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Add(login, password, name, email string, isSuperuser bool) error {

	_, err := env.Connect().Exec(`SELECT users.add(
			$1,
			$2,
			$3,
			$4,
			$5
		)`,
		login,
		password,
		name,
		email,
		isSuperuser,
	)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not add user: %v", err)
		}

		return err
	}

	return nil
}
