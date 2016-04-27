package auth

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/env"
	"github.com/postgres-ci/app-server/src/tools/db"

	"fmt"
	"time"
)

const (
	CookieName = "sid"
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

		if !db.NotFound(err) {

			log.Errorf("Error when receiving auth user: %v", err)
		}

		return nil, err
	}

	return &user, nil
}

func Login(login, password string) (string, error) {

	var sessionID string

	if err := env.Connect().Get(&sessionID, `SELECT session_id FROM auth.login($1, $2)`, login, password); err != nil {

		if db.NotFound(err) || db.InvalidPassword(err) {

			return "", fmt.Errorf("INVALID_LOGIN_OR_PASSWORD")
		}

		log.Errorf("Error when execute 'login' query: %v", err)

		return "", err
	}

	return sessionID, nil
}
