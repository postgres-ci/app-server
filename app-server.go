package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/postgres-ci/app-server/src/app"
	"github.com/postgres-ci/app-server/src/common"
	"github.com/postgres-ci/app-server/src/tools/render"

	"flag"
	"fmt"
	"os"
)

var (
	debug        bool
	pathToConfig string
)

const usage = `
Postgres-CI app-server

Usage:
    -c /path/to/config.yaml (if not set, then app-server will use environment variables)
    -debug (enable debug mode)

Environment variables:

    ADDRESS   - bind address ip:port
    TEMPLATES - path to templates dir
    LOG_LEVEL - one of: info/warning/error

    == PostgreSQL server credentials

    DB_HOST
    DB_PORT
    DB_USERNAME
    DB_PASSWORD
    DB_DATABASE 
`

func main() {

	flag.BoolVar(&debug, "debug", false, "")
	flag.StringVar(&pathToConfig, "c", "", "")

	flag.Usage = func() {

		fmt.Println(usage)

		os.Exit(0)
	}

	flag.Parse()

	if log.IsTerminal() {

		log.SetFormatter(&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05 MST",
		})

	} else {

		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05 MST",
		})
	}

	config, err := common.ReadConfig(pathToConfig)

	if err != nil {

		log.Fatalf("Error reading configuration file: %v", err)
	}

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(config.LogLevel())
	}

	if err := render.Init(config.Templates); err != nil {

		log.Fatalf("Error when loading templates: %v", err)
	}

	app := app.New(config)

	app.Run()
}
