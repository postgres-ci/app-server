package webhooks

import (
	"github.com/postgres-ci/app-server/src/env"
	"github.com/postgres-ci/hooks/git"

	"strings"
)

func Commit(token, ref string, commit git.Commit) error {

	var branch string

	if parts := strings.Split(ref, "/"); len(parts) > 0 {

		branch = parts[len(parts)-1]
	}

	_, err := env.Connect().Exec(`SELECT hook.commit($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		token,
		branch,
		commit.ID,
		commit.Message,
		commit.CommittedAt,
		commit.Committer.Name,
		commit.Committer.Email,
		commit.Author.Name,
		commit.Author.Email,
	)

	if err != nil {

		return err
	}

	return nil
}
