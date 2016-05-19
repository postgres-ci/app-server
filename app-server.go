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
    -c /path/to/config.yaml (if not setted app will use environment variables)
    -debug (enable debug mode)

Environment variables:

    APP_ADDRESS (example: 127.0.0.1:8888)
    APP_TEMPLATES (example: /opt/postgres-ci/app-server/templates/)
    APP_LOG_LEVEL (one of: info/warning/error)
    DB_HOST (example: 10.20.11.42)
    DB_PORT (example: 5432)
    DB_USERNAME (example: postgres_ci)
    DB_PASSWORD (example: PcSd23@@a)
    DB_DATABASE (example: postgres_ci)
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
