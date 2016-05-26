package github

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/app/models/webhooks/common"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"
)

func Push(push PushEvent) error {

	commits := make(common.Commits, 0, len(push.Commits))

	for _, c := range push.Commits {

		commits = append(commits, common.Commit{
			SHA:            c.ID,
			Message:        c.Message,
			CommittedAt:    c.Timestamp,
			CommitterName:  c.Committer.Name,
			CommitterEmail: c.Committer.Email,
			AuthorName:     c.Author.Name,
			AuthorEmail:    c.Author.Email,
		})
	}

	_, err := env.Connect().Exec("SELECT hook.github_push($1, $2, $3)", push.Repository.FullName, common.RefToBranch(push.Ref), commits)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not add GitHub push: %v", err)
		}

		return err
	}

	return nil
}
