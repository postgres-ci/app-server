package common

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ReadConfig(path string) (Config, error) {

	var config Config

	if path == "" {

		config.setFromEnv()

		return config, nil
	}

	if _, err := os.Open(path); err != nil {

		if os.IsNotExist(err) {

			return config, fmt.Errorf("No such configuration file '%s'", path)
		}

		return config, fmt.Errorf("Could not open configuration file '%s': %v", path, err)
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {

		return config, nil
	}

	if err := yaml.Unmarshal(data, &config); err != nil {

		return config, err
	}

	return config, nil
}

type Config struct {
	Address   string  `yaml:"address"`
	Templates string  `yaml:"templates"`
	Connect   Connect `yaml:"connect"`
	Loglevel  string  `yaml:"loglevel"`
}

type Connect struct {
	Host     string `yaml:"host"`
	Port     uint32 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (c *Config) setFromEnv() {

	var port uint32 = 5432

	if value, err := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32); err == nil {

		port = uint32(value)
	}

	c.Address = os.Getenv("APP_ADDRESS")
	c.Templates = os.Getenv("APP_TEMPLATES")
	c.Loglevel = os.Getenv("APP_LOG_LEVEL")
	c.Connect = Connect{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}

func (c *Connect) DSN() string {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Username, c.Password, c.Database)

	if c.Port != 0 {

		dsn += fmt.Sprintf(" port=%d", c.Port)
	}

	return dsn
}

func (c *Config) LogLevel() log.Level {

	switch c.Loglevel {
	case "info":
		return log.InfoLevel
	case "warning":
		return log.WarnLevel
	}

	return log.ErrorLevel
}
