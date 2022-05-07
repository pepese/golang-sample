package app

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	AppName     string `envconfig:"APP_NAME" default:"golang-sample"`
	AppVersion  string `envconfig:"APP_VERSION" default:"undefined"`
	RdbType     string `envconfig:"RDB_TYPE" default:"postgres"`
	RdbUser     string `envconfig:"RDB_USER" default:"testuser"`
	RdbPassword string `envconfig:"RDB_PASSWORD" default:"testpass"`
	RdbProtocol string `envconfig:"RDB_PROTOCOL" default:"tcp"`
	RdbHost     string `envconfig:"RDB_HOST" default:"127.0.0.1"`
	RdbPort     string `envconfig:"RDB_PORT" default:"5432"`
	RdbName     string `envconfig:"RDB_NAME" default:"testdb"`
	TimeZone    string `envconfig:"TZ" default:"Asia/Tokyo"`
}

var (
	conf  *config
	cOnce sync.Once
)

func Config() *config {
	initConfig()
	return conf
}

func initConfig() {
	cOnce.Do(func() {
		conf = &config{}
		if err := envconfig.Process("", conf); err != nil {
			panic("envconfig error")
		}
	})
}
