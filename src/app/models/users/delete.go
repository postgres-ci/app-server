package users

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/env"
)

func Delete(userID int32) error {

	if _, err := env.Connect().Exec("SELECT users.delete($1)", userID); err != nil {

		log.Errorf("Could not delete user %d, err: %v", userID, err)

		return err
	}

	return nil
}
