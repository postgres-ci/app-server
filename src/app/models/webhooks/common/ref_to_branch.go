package common

import (
	"strings"
)

func RefToBranch(ref string) string {

	var branch string

	if parts := strings.Split(ref, "/"); len(parts) > 0 {

		branch = parts[len(parts)-1]
	}

	return branch
}
