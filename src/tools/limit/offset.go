package limit

import (
	"github.com/postgres-ci/http200ok"

	"strconv"
)

func Offset(c *http200ok.Context, limit int32) int32 {

	var currentPage int32

	if value, err := strconv.ParseInt(c.Request.URL.Query().Get("p"), 10, 32); err == nil {

		currentPage = int32(value)
	}

	if currentPage <= 1 {

		return 0
	}

	if offset := (limit * (currentPage - 1)); offset > 0 {

		return offset
	}

	return 0
}
