package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/go-city/command"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "go-city",
		Usage: "go city",
		Commands: []*cli.Command{
			{
				Name:        "sync",
				Aliases:     []string{"s"},
				Description: "同步省市区到数据库，数据来源于京东",
				Flags:       command.Flags,
				Action:      command.Run,
			},
			{
				Name:        "sync-tjj",
				Aliases:     []string{"tjj"},
				Description: "同步省市区到数据库，数据来源于国家统计局",
				Flags:       command.Flags,
				Action:      command.Exec,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Error(err.Error())
	}
}
