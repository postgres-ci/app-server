package project

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/env"

	"encoding/json"
	"fmt"
	"time"
)

type user struct {
	Id   int32  `json:"user_id"`
	Name string `json:"user_name"`
}

type users []user

func (u *users) Scan(src interface{}) error {

	var source []byte

	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return fmt.Errorf("Incompatible type for users")
	}

	return json.Unmarshal(source, u)
}

type Project struct {
	ID             int32     `db:"project_id"       json:"project_id"`
	Name           string    `db:"project_name"     json:"project_name"`
	Token          string    `db:"project_token"    json:"project_token"`
	URL            string    `db:"repository_url"   json:"repository_url"`
	OwnerID        int32     `db:"project_owner_id" json:"project_owner_id"`
	PossibleOwners users     `db:"possible_owners"  json:"possible_owners"`
	GitHubSecret   string    `db:"github_secret"    json:"github_secret"`
	CreatedAt      time.Time `db:"created_at"       json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"       json:"updated_at"`
}

func Get(projectID int32) (Project, error) {

	var project Project

	err := env.Connect().Get(&project, `
		SELECT 
			project_id,
			project_name,
			project_token,
			repository_url,
			project_owner_id,
			possible_owners,
			github_secret,
			created_at,
			updated_at
		FROM project.get($1)
		`,
		projectID,
	)

	if err != nil {

		err := errors.Wrap(err)

		if err.(*errors.Error).IsFatal() {

			log.Errorf("Could not fetch project: %v", err)
		}

		return project, err
	}

	return project, nil
}
