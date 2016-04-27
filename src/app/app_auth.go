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

	app.Get("/login/", loginHandler)
	app.Post("/login/", loginHandler)
}

func loginHandler(c *http200ok.Context) {

	if c.Request.Method == "POST" {

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

		return
	}

	render.HTML(c, "login.html", nil)
}
