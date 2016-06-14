package notification

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func UpdateMethod(userID int32, method, textID string) error {

	if _, err := env.Connect().Exec("SELECT notification.update_method($1, $2, $3)", userID, method, textID); err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not update notification method: %v", err)
		}

		return err
	}

	return nil
}
