package auth

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Login(login, password string) (string, error) {

	var sessionID string

	if err := env.Connect().Get(&sessionID, `SELECT session_id FROM auth.login($1, $2)`, login, password); err != nil {

		if !(errors.IsNotFound(err) || errors.IsInvalidPassword(err)) {

			log.Errorf("Error when execute 'login' query: %v", err)
		}

		return "", err
	}

	return sessionID, nil
}
