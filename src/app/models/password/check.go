package password

import (
	"github.com/postgres-ci/app-server/src/env"
)

func Check(userID int32, password string) error {

	if _, err := env.Connect().Exec("SELECT password.check($1, $2)", userID, password); err != nil {

		return err
	}

	return nil
}
