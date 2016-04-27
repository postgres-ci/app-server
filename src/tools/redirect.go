package tools

import (
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func Redirect(c *http200ok.Context, urlStr string) {

	c.Response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Response.Header().Set("Expires", "0")

	http.Redirect(c.Response, c.Request, urlStr, http.StatusFound)
}
