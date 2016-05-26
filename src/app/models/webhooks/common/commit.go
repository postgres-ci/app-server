package common

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Commit struct {
	SHA            string    `json:"commit_sha"`
	Message        string    `json:"commit_message"`
	CommittedAt    time.Time `json:"committed_at"`
	CommitterName  string    `json:"committer_name"`
	CommitterEmail string    `json:"committer_email"`
	AuthorName     string    `json:"author_name"`
	AuthorEmail    string    `json:"author_email"`
}

type Commits []Commit

func (c Commits) Value() (driver.Value, error) {
	return json.Marshal(c)
}
