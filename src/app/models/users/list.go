package users

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/env"

	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID          int32     `db:"user_id"      json:"user_id"`
	Name        string    `db:"user_name"    json:"user_name"`
	Login       string    `db:"user_login"   json:"user_login"`
	Email       string    `db:"user_email"   json:"user_email"`
	IsSuperuser bool      `db:"is_superuser" json:"is_superuser"`
	CreatedAt   time.Time `db:"created_at"   json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"   json:"updated_at"`
}

type Users []User

func (u *Users) Scan(src interface{}) error {

	var source []byte

	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return fmt.Errorf("Incompatible type for Users")
	}

	return json.Unmarshal(source, u)
}

type list struct {
	Total int32 `db:"total"`
	Users Users `db:"users"`
}

func List(limit, offset int32, query string) (*list, error) {

	var list list

	err := env.Connect().Get(&list, `SELECT total, users FROM users.list($1, $2, $3)`, limit, offset, strings.TrimSpace(query))

	if err != nil {

		log.Errorf("Error when fetching users: %v", err)

		return nil, err
	}

	return &list, nil
}
