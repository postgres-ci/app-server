package middleware

import (
	"github.com/postgres-ci/app-server/src/app/auth"
	"github.com/postgres-ci/app-server/src/tools"
	"github.com/postgres-ci/http200ok"

	"strings"
)

func Authentication(c *http200ok.Context) {

	if strings.HasPrefix(c.Request.URL.Path, "/webhooks/") || c.Request.URL.Path == "/login/" {

		return
	}

	if cookie, err := c.Request.Cookie(auth.CookieName); err == nil {

		if user, err := auth.GetUser(cookie.Value); err == nil {

			c.Set("CurrentUser", user)

			return
		}
	}

	c.Set("CurrentUser", &auth.User{})

	tools.Redirect(c, "/login/")

	c.Stop()
}
