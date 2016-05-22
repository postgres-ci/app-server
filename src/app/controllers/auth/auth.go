package auth

import (
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/common/errors"
	"github.com/postgres-ci/app-server/src/tools"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func Route(server *http200ok.Server) {

	server.Post("/login/", loginHandler)

	server.Get("/login/", func(c *http200ok.Context) {

		render.HTML(c, "login.html", nil)
	})

	server.Get("/logout/", func(c *http200ok.Context) {

		if cookie, err := c.Request.Cookie(auth.CookieName); err == nil {

			auth.Logout(cookie.Value)
		}

		http.SetCookie(c.Response, &http.Cookie{
			Name:     auth.CookieName,
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		})

		tools.Redirect(c, "/")
	})
}

func loginHandler(c *http200ok.Context) {

	var (
		login    = c.Request.PostFormValue("login")
		password = c.Request.PostFormValue("password")
	)

	sessioID, err := auth.Login(login, password)

	if err != nil {

		var _error = err.Error()

		if errors.IsNotFound(err) || errors.IsInvalidPassword(err) {

			_error = "INVALID_LOGIN_OR_PASSWORD"
		}

		render.HTML(c, "login.html", render.Context{

			"error": _error,
		})

		return
	}

	http.SetCookie(c.Response, &http.Cookie{
		Name:     auth.CookieName,
		Value:    sessioID,
		Path:     "/",
		HttpOnly: true,
	})

	tools.Redirect(c, "/")
}
