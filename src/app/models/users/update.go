package users

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Update(userID int32, name, email string, isSuperuser bool) error {

	_, err := env.Connect().Exec(`SELECT users.update(
			$1,
			$2,
			$3,
			$4
		)`,
		userID,
		name,
		email,
		isSuperuser,
	)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not update user %d, err: %v", userID, err)
		}

		return err
	}

	return nil
}
