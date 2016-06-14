package notification

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

type method struct {
	Method string `db:"method"  json:"method"`
	TextID string `db:"text_id" json:"text_id"`
	IntID  int64  `db:"int_id"  json:"int_id"`
}

func GetMethod(userID int32) (*method, error) {

	var result method

	err := env.Connect().Get(&result, `SELECT method, text_id, int_id FROM notification.get_method($1)`, userID)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not fetch user notification method: %v", err)
		}

		return nil, err
	}

	return &result, nil
}
