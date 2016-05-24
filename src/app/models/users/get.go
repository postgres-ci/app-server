package users

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Get(userID int32) (User, error) {

	var user User

	err := env.Connect().Get(&user, `
		SELECT 
			user_id,
			user_name,
			user_login,
			user_email,
			is_superuser,
			created_at,
			updated_at
		FROM users.get($1)
		`,
		userID,
	)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not fetch user: %v", err)
		}

		return user, err
	}

	return user, nil
}
