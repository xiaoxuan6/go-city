package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/go-city/command"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "go-city",
		Usage:       "go city",
		Description: figure.NewFigure("Go City", "", true).String() + "同步远程数据到数据库",
		Commands: []*cli.Command{
			{
				Name:        "sync",
				Aliases:     []string{"s"},
				Description: figure.NewFigure("Go City Sync", "", true).String() + "同步省市区到数据库，数据来源于京东",
				Flags:       command.Flags,
				Action:      command.Run,
			},
			{
				Name:        "sync-tjj",
				Aliases:     []string{"tjj"},
				Description: figure.NewFigure("Go City Sync-Tjj", "", true).String() + "同步省市区到数据库，数据来源于国家统计局",
				Flags:       command.Flags,
				Action:      command.Exec,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Error(err.Error())
	}
}
