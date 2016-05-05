package app

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/app/auth"
	"github.com/postgres-ci/app-server/src/tools"
	"github.com/postgres-ci/app-server/src/tools/render"
	"github.com/postgres-ci/http200ok"

	"net/http"
)

func (app *app) auth() {

	app.Post("/login/", loginHandler)

	app.Get("/login/", func(c *http200ok.Context) {

		render.HTML(c, "login.html", nil)
	})

	app.Get("/logout/", func(c *http200ok.Context) {

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

	log.Debugf("login: %s, password: %s -> %s, %v", login, password, sessioID, err)

	if err != nil {

		render.HTML(c, "login.html", render.Context{

			"error": err.Error(),
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
