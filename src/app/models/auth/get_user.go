package auth

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"

	"time"
)

type User struct {
	ID          int32     `db:"user_id"`
	Name        string    `db:"user_name"`
	Login       string    `db:"user_login"`
	Email       string    `db:"user_email"`
	IsSuperuser bool      `db:"is_superuser"`
	CreatedAt   time.Time `db:"created_at"`
}

func GetUser(sessionID string) (*User, error) {

	var user User

	err := env.Connect().Get(&user, `
		SELECT 
			user_id,
			user_name,
			user_login,
			user_email,
			is_superuser,
			created_at
		FROM auth.get_user($1)
	`, sessionID)

	if err != nil {

		if !errors.IsNotFound(err) {

			log.Errorf("Error when receiving auth user: %v", err)
		}

		return nil, err
	}

	return &user, nil
}
