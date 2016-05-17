package webhooks

import (
	"github.com/postgres-ci/app-server/src/app/models/webhooks/common"
	"github.com/postgres-ci/app-server/src/env"
	"github.com/postgres-ci/hooks/git"

	"database/sql/driver"
	"encoding/json"
	"time"
)

type commit struct {
	SHA            string    `json:"commit_sha"`
	Message        string    `json:"commit_message"`
	CommittedAt    time.Time `json:"committed_at"`
	CommitterName  string    `json:"committer_name"`
	CommitterEmail string    `json:"committer_email"`
	AuthorName     string    `json:"author_name"`
	AuthorEmail    string    `json:"author_email"`
}

type commits []commit

func (c commits) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func Push(token string, push git.Push) error {

	commits := make(commits, 0, len(push.Commits))

	for _, c := range push.Commits {

		commits = append(commits, commit{
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

		return err
	}

	return nil
}
