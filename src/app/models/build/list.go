package build

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/env"

	"encoding/json"
	"fmt"
	"time"
)

type Item struct {
	BuildID       int32      `json:"build_id"`
	Status        string     `json:"status"`
	Error         string     `json:"error"`
	CreatedAt     time.Time  `json:"created_at"`
	StartedAt     *time.Time `json:"started_at"`
	FinishedAt    *time.Time `json:"finished_at"`
	CommitSHA     string     `json:"commit_sha"`
	CommitMessage string     `json:"commit_message"`
	Branch        string     `json:"branch"`
	BranchID      int32      `json:"branch_id"`
	ProjectID     int32      `json:"project_id"`
}

type Items []Item

func (i *Items) Scan(src interface{}) error {

	var source []byte

	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return fmt.Errorf("Incompatible type for Items")
	}

	return json.Unmarshal(source, i)
}

type Branch struct {
	ID   int32  `json:"branch_id"`
	Name string `json:"branch"`
}

type Branches []Branch

func (b *Branches) Scan(src interface{}) error {

	var source []byte

	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return fmt.Errorf("Incompatible type for Branches")
	}

	return json.Unmarshal(source, b)
}

type list struct {
	Branches Branches `db:"branches"`
	Total    int32    `db:"total"`
	Items    Items    `db:"items"`
}

func List(projectID, branchID, limit, offset int32) (*list, error) {

	var list list

	err := env.Connect().Get(&list, `SELECT total, branches, items FROM build.list($1, $2, $3, $4)`, projectID, branchID, limit, offset)

	if err != nil {

		log.Errorf("Error when fetching list of builds: %v", err)

		return nil, err
	}

	return &list, nil
}
