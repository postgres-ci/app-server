package webhooks

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/app/models/webhooks/common"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
	"github.com/postgres-ci/hooks/git"
)

func Push(token string, push git.Push) error {

	commits := make(common.Commits, 0, len(push.Commits))

	for _, c := range push.Commits {

		commits = append(commits, common.Commit{
			SHA:            c.ID,
			Message:        c.Message,
			CommittedAt:    c.CommittedAt,
			CommitterName:  c.Committer.Name,
			CommitterEmail: c.Committer.Email,
			AuthorName:     c.Author.Name,
			AuthorEmail:    c.Author.Email,
		})
	}

	_, err := env.Connect().Exec("SELECT hook.push($1, $2, $3)", token, common.RefToBranch(push.Ref), commits)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not add push: %v", err)
		}

		return err
	}

	return nil
}
