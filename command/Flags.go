package command

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "driver",
		Usage:    "db driver，Support：sqlite、database、memory",
		Required: false,
		Value:    "sqlite",
		Aliases:  []string{"d"},
		EnvVars:  []string{"GO_CITY_DRIVER"},
	},
	&cli.StringFlag{
		Name:     "db",
		Usage:    "database file",
		Required: false,
		Value:    "./sqlite.db",
		EnvVars:  []string{"GO_CITY_DB"},
		Action: func(context *cli.Context, s string) error {
			if strings.HasSuffix(s, ".db") == false {
				return errors.New(fmt.Sprintf("db 文件 %s 后缀错误，必须是'.db'", s))
			}
			return nil
		},
	},
	&cli.StringFlag{
		Name:     "host",
		Usage:    "database host",
		Required: false,
		Value:    "127.0.0.1",
		EnvVars:  []string{"GO_CITY_DB_HOST"},
	},
	&cli.StringFlag{
		Name:     "port",
		Usage:    "database port",
		Required: false,
		Value:    "3306",
		EnvVars:  []string{"GO_CITY_DB_PORT"},
	},
	&cli.StringFlag{
		Name:     "username",
		Usage:    "database username",
		Required: false,
		Value:    "root",
		EnvVars:  []string{"GO_CITY_DB_USERNAME"},
	},
	&cli.StringFlag{
		Name:     "password",
		Usage:    "database password",
		Required: false,
		Value:    "root",
		EnvVars:  []string{"GO_CITY_DB_PASSWORD"},
	},
	&cli.StringFlag{
		Name:     "dbname",
		Usage:    "database dbname",
		Required: false,
		Value:    "",
		EnvVars:  []string{"GO_CITY_DB_NAME"},
	},
	&cli.StringFlag{
		Name:     "table",
		Usage:    "database table name",
		Required: false,
		Value:    "cities",
		EnvVars:  []string{"GO_CITY_DB_TABLE_NAME"},
	},
}
