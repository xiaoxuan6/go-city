package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/xiaoxuan6/go-city/command"
	"os"
)

var Version string

func main() {
	app := &cli.App{
		Name:        "go-city",
		Usage:       "go city",
		Description: figure.NewFigure("Go City", "", true).String() + "同步远程数据到数据库",
		Commands: []*cli.Command{
			{
				Name:        "sync",
				Aliases:     []string{"s"},
				Usage:       "sync data to database, data source is jd",
				Description: figure.NewFigure("Go City Sync", "", true).String() + "同步省市区到数据库，数据来源于京东",
				Flags:       command.Flags,
				Action:      command.Run,
			},
			{
				Name:        "sync-tjj",
				Aliases:     []string{"tjj"},
				Usage:       "sync data to database, data source is 统计局",
				Description: figure.NewFigure("Go City Sync-Tjj", "", true).String() + "同步省市区到数据库，数据来源于国家统计局",
				Flags:       command.Flags,
				Action:      command.Exec,
			},
			{
				Name:        "version",
				Aliases:     []string{"v"},
				Usage:       "show go city version",
				Description: figure.NewFigure("Go City Version", "", true).String() + "显示版本号",
				Action: func(context *cli.Context) error {
					logrus.Info("go-city version: ", Version)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Error(err.Error())
	}
}
