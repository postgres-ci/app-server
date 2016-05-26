package render

import (
	log "github.com/Sirupsen/logrus"
	"github.com/flosch/pongo2"
	_ "github.com/postgres-ci/app-server/src/tools/render/filters"
	"github.com/postgres-ci/http200ok"

	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type Context map[string]interface{}

var templates = make(map[string]*pongo2.Template)

func Init(root string) error {

	if _, err := os.Open(root); err != nil {

		return err
	}

	return filepath.Walk(root, func(path string, fi os.FileInfo, _ error) error {

		if fi == nil || fi.IsDir() {

			return nil
		}

		if name, err := filepath.Rel(root, path); err == nil {

			template, err := pongo2.FromFile(path)

			if err != nil {

				return err
			}

			templates[name] = template

		} else {

			return err
		}

		return nil
	})
}

func HTML(c *http200ok.Context, name string, context Context) error {

	if currentUser := c.Get("CurrentUser"); currentUser != nil {

		if context == nil {

			context = make(Context)
		}

		context["CurrentUser"] = currentUser
	}

	if template, found := templates[name]; found {

		if err := template.ExecuteWriterUnbuffered(pongo2.Context(context), c.Response); err != nil {

			log.Errorf("Render: %v", err)

			return err
		}

		return nil
	}

	log.Errorf("Template '%s' not found", name)

	return fmt.Errorf("template '%s' not found", name)
}

func NotFound(c *http200ok.Context) {

	http.NotFound(c.Response, c.Request)
}

func InternalServerError(c *http200ok.Context, err error) {

	c.Response.WriteHeader(http.StatusInternalServerError)
}

func JSON(c *http200ok.Context, v interface{}) error {

	c.Response.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(c.Response).Encode(v)
}

func JSONok(c *http200ok.Context) {

	JSON(c, struct {
		Success bool `json:"success"`
	}{
		Success: true,
	})
}
func JSONError(c *http200ok.Context, code int, format string, a ...interface{}) error {

	c.Response.Header().Add("Content-Type", "application/json")
	c.Response.Header().Set("X-Content-Type-Options", "nosniff")
	c.Response.WriteHeader(code)

	return json.NewEncoder(c.Response).Encode(struct {
		Success bool   `json:"success"`
		Code    int    `json:"code"`
		Error   string `json:"error"`
	}{
		Success: false,
		Code:    code,
		Error:   fmt.Sprintf(format, a...),
	})
}
