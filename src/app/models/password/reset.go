package password

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Reset(userID int32, password string) error {

	if _, err := env.Connect().Exec("SELECT password.reset($1, $2)", userID, password); err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not reset password: %v", err)
		}

		return err
	}

	return nil
}
