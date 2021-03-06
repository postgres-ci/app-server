package password

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Check(userID int32, password string) error {

	if _, err := env.Connect().Exec("SELECT password.check($1, $2)", userID, password); err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not check password: %v", err)
		}

		return err
	}

	return nil
}
