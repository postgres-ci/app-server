package project

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/env"

	"time"
)

type ListItem struct {
	ProjectID      int32     `db:"project_id"`
	ProjectName    string    `db:"project_name"`
	ProjectToken   string    `db:"project_token"`
	ProjectOwnerID int32     `db:"project_owner_id"`
	UserEmail      string    `db:"user_email"`
	UserName       string    `db:"user_name"`
	Status         string    `db:"status"`
	CommitSHA      string    `db:"commit_sha"`
	LastBuildID    int32     `db:"last_build_id"`
	StartedAt      time.Time `db:"started_at"`
	FinishedAt     time.Time `db:"finished_at"`
}

func List() ([]ListItem, error) {

	var result []ListItem

	err := env.Connect().Select(&result, `
		SELECT 
			project_id,
			project_name,
			project_token,
			project_owner_id,
			user_email,
			user_name,
			COALESCE(status::text, 'n/a') AS status,
			COALESCE(commit_sha, '')      AS commit_sha,
			COALESCE(last_build_id, 0)    AS last_build_id,
			COALESCE(started_at,  '0001-01-01 00:00:00+00') AS started_at,
			COALESCE(finished_at, '0001-01-01 00:00:00+00') AS finished_at 
		FROM project.list()
	`)

	if err != nil {

		log.Errorf("Error when fetching list of projects: %v", err)

		return nil, err
	}

	return result, nil
}
