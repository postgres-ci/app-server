package password

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Change(userID int32, currentPassword, newPassword string) error {

	if _, err := env.Connect().Exec("SELECT password.change($1, $2, $3)", userID, currentPassword, newPassword); err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not change password: %v", err)
		}

		return err
	}

	return nil
}
