package webhooks

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/app/models/webhooks/common"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
	"github.com/postgres-ci/hooks/git"
)

func Commit(token, ref string, commit git.Commit) error {

	_, err := env.Connect().Exec(`SELECT hook.commit($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		token,
		common.RefToBranch(ref),
		commit.ID,
		commit.Message,
		commit.CommittedAt,
		commit.Committer.Name,
		commit.Committer.Email,
		commit.Author.Name,
		commit.Author.Email,
	)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not add commit: %v", err)
		}

		return err
	}

	return nil
}
