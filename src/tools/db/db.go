package db

import (
	"database/sql"
	"github.com/lib/pq"
)

func NotFound(err error) bool {

	if err == sql.ErrNoRows {

		return true
	}

	if err, ok := err.(*pq.Error); ok && err.Code.Name() == "no_data_found" {

		return true
	}

	return false
}

func InvalidPassword(err error) bool {

	if err, ok := err.(*pq.Error); ok && err.Code.Name() == "invalid_password" {

		return true
	}

	return false
}
