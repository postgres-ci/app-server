package github

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Secret(fullName string) (string, error) {

	var secret string

	if err := env.Connect().Get(&secret, "SELECT secret FROM project.github_secret($1)", fullName); err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not fetch GitHub secret: %v", err)
		}

		return "", err
	}

	return secret, nil
}
