package build

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/env"

	"encoding/json"
	"fmt"
	"time"
)

type build struct {
	ProjectID      int32     `db:"project_id"`
	ProjectName    string    `db:"project_name"`
	BranchID       int32     `db:"branch_id"`
	Branch         string    `db:"branch_name"`
	Config         string    `db:"config"`
	Status         string    `db:"status"`
	Error          string    `db:"error"`
	CommitSHA      string    `db:"commit_sha"`
	CommitMessage  string    `db:"commit_message"`
	CommittedAt    time.Time `db:"committed_at"`
	CommitterName  string    `db:"committer_name"`
	CommitterEmail string    `db:"committer_email"`
	AuthorName     string    `db:"author_name"`
	AuthorEmail    string    `db:"author_email"`
	Parts          Parts     `db:"parts"`
}

type Part struct {
	PartID     int32     `json:"part_id"`
	Image      string    `json:"image"`
	Version    string    `json:"version"`
	Output     string    `json:"output"`
	Success    bool      `json:"success"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	Tests      []Test    `json:"tests"`
}

type Parts []Part

func (p *Parts) Scan(src interface{}) error {

	var source []byte

	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return fmt.Errorf("Incompatible type for Parts")
	}

	return json.Unmarshal(source, p)
}

type Test struct {
	Function string  `json:"function"`
	Errors   []Error `json:"errors"`
	Duration float64 `json:"duration"`
}

type Error struct {
	Message string `json:"message"`
	Comment string `json:"comment"`
}

func View(buildID int32) (*build, error) {

	var build build

	err := env.Connect().Get(&build, `SELECT * FROM build.view($1)`, buildID)

	if err != nil {

		log.Errorf("Error when fetching build: %v", err)

		return nil, err
	}

	return &build, nil
}
