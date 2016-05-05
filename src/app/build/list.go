package build

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/env"

	"encoding/json"
	"fmt"
	"time"
)

type Item struct {
	BuildID     int32      `json:"build_id"`
	Status      string     `json:"status"`
	Error       string     `json:"error"`
	CreatedAt   time.Time  `json:"created_at"`
	StartedAt   *time.Time `json:"started_at"`
	FinishedAt  *time.Time `json:"finished_at"`
	CommitSHA   string     `json:"commit_sha"`
	Branch      string     `json:"branch"`
	ProjectName string     `json:"project_name"`
	ProjectID   int32      `json:"project_id"`
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

func List(limit, offset uint32) (int64, []Item, error) {

	var result struct {
		Total int64 `db:"total"`
		Items Items `db:"items"`
	}

	err := env.Connect().Get(&result, `SELECT total, items FROM build.list($1, $2)`, limit, offset)

	if err != nil {

		log.Errorf("%v", err)

		return 0, nil, err
	}

	return result.Total, result.Items, nil
}
