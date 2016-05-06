package github

import (
	"time"
)

type PushEvent struct {
	Ref        string   `json:"ref"`
	Before     string   `json:"before"`
	After      string   `json:"after"`
	Commits    []Commit `json:"commits"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

type Committer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Commit struct {
	ID        string    `json:"id"`
	Author    Committer `json:"author"`
	Committer Committer `json:"committer"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
