package command

import "github.com/urfave/cli/v2"

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "driver",
		Usage:    "db driver，Support：sqlite、database、memory",
		Required: false,
		Value:    "sqlite",
		Aliases:  []string{"d"},
		EnvVars:  []string{"GO_CITY_DRIVER"},
	},
}
