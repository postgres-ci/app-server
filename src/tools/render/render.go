package render

import (
	log "github.com/Sirupsen/logrus"
	"github.com/flosch/pongo2"
	"github.com/postgres-ci/http200ok"

	"encoding/json"
	"fmt"
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

		return template.ExecuteWriterUnbuffered(pongo2.Context(context), c.Response)
	}

	log.Errorf("Template '%s' not found", name)

	return fmt.Errorf("template '%s' not found", name)
}

func JSON(c *http200ok.Context, v interface{}) error {

	c.Response.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(c.Response).Encode(v)
}

func JSONError(c *http200ok.Context, code int, err interface{}) error {

	c.Response.Header().Add("Content-Type", "application/json")
	c.Response.Header().Set("X-Content-Type-Options", "nosniff")
	c.Response.WriteHeader(code)

	return json.NewEncoder(c.Response).Encode(struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Error   interface{} `json:"error"`
	}{
		Success: false,
		Code:    code,
		Error:   err,
	})
}