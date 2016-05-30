package app

import (
	"github.com/Sirupsen/logrus"
	"github.com/erikstmartin/go-testdb"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/postgres-ci/app-server/src/app/middleware"
	"github.com/postgres-ci/app-server/src/app/models/auth"
	"github.com/postgres-ci/app-server/src/env"
	"github.com/postgres-ci/http200ok"
	"github.com/stretchr/testify/assert"

	"bytes"
	"database/sql/driver"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func CreateApp() *app {

	logrus.SetOutput(&bytes.Buffer{})

	connect, err := sqlx.Connect("testdb", "")

	if err != nil {

		log.Fatalf("Could not connect to database server: %v", err)
	}

	env.SetConnect(connect)

	app := &app{
		Server: http200ok.New(),
	}

	app.Use(middleware.Authentication)
	app.route()

	return app
}

func TestAuthentication(t *testing.T) {

	server := httptest.NewServer(CreateApp())
	testdb.EnableTimeParsing(true)

	defer testdb.Reset()

	sql := `
		SELECT
			user_id,
			user_name,
			user_login,
			user_email,
			is_superuser,
			created_at
		FROM auth.get_user($1)
		`

	testdb.StubQueryError(sql, &pq.Error{Code: "P0002"})

	for _, url := range []string{
		"/",
		"/project-1/get/",
		"/project-1/builds/",
		"/project-1/builds/branch-1/",
		"/project-1/build-1/",
		"/project/possible-owners/",
		"/users/",
		"/users/get/1/",
	} {

		req, _ := http.NewRequest("GET", server.URL+url, nil)
		req.Header.Set("Cookie", fmt.Sprintf("%s=cookie", auth.CookieName))

		if response, err := (&http.Transport{}).RoundTrip(req); assert.NoError(t, err) {

			if assert.Equal(t, http.StatusFound, response.StatusCode) {
				assert.Equal(t, "/login/", response.Header.Get("Location"))
			}
		}
	}

	for _, url := range []string{
		"/project/add/",
		"/project/update/1/",
		"/project/delete/1/",
		"/password/change/",
		"/password/reset/1/",
		"/users/add/",
		"/users/update/1/",
		"/users/delete/1/",
	} {

		req, _ := http.NewRequest("POST", server.URL+url, nil)
		req.Header.Set("Cookie", fmt.Sprintf("%s=cookie", auth.CookieName))

		if response, err := (&http.Transport{}).RoundTrip(req); assert.NoError(t, err) {

			if assert.Equal(t, http.StatusFound, response.StatusCode) {
				assert.Equal(t, "/login/", response.Header.Get("Location"))
			}
		}
	}

	req, _ := http.NewRequest("GET", server.URL+"/users/", nil)
	req.Header.Set("Cookie", fmt.Sprintf("%s=cookie", auth.CookieName))

	testdb.SetQueryWithArgsFunc(func(query string, args []driver.Value) (result driver.Rows, err error) {

		switch true {

		case strings.Contains(query, "auth.get_user"):
			return testdb.RowsFromCSVString(
				[]string{"user_id", "user_name", "user_login", "user_email", "is_superuser", "created_at"},
				"1,user,login,email@email.com,true,2016-05-30T11:02:41+03:00",
			), nil

		case strings.Contains(query, "users.list"):

			return testdb.RowsFromCSVString(
				[]string{"total", "users"},
				"42,[]",
			), nil
		}

		return nil, fmt.Errorf("SQL_ERROR")
	})

	if response, err := (&http.Transport{}).RoundTrip(req); assert.NoError(t, err) {

		assert.Equal(t, http.StatusOK, response.StatusCode)
	}
}
