package params

import (
	"github.com/postgres-ci/http200ok"
	"strconv"
)

func ToInt32(c *http200ok.Context, key string) int32 {

	if v, err := strconv.ParseInt(c.RequestParam(key), 10, 32); err == nil {

		return int32(v)
	}

	return 0
}
