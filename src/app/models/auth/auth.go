package auth

import (
	"github.com/postgres-ci/app-server/src/env"
)

const (
	CookieName = "sid"
)

func Logout(sessionID string) {

	env.Connect().Exec("SELECT auth.logout($1)", sessionID)
}
