package common

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestRefToBranch(t *testing.T) {

	assets := map[string]string{
		"refs/heads/master":  "master",
		"refs/heads/changes": "changes",
		"refs/heads/v1":      "v1",
		"branch":             "branch",
	}

	for ref, expected := range assets {

		assert.Equal(t, expected, RefToBranch(ref))
	}
}
