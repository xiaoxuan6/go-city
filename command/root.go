package command

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func base(c *cli.Context) error {
	force := c.Bool("force")
	driver := c.String("driver")
	switch driver {
	case "sqlite":

		if force {
			filepath := ""
			db := c.String("db")
			if strings.HasPrefix(db, "./") {
				path, _ := os.Getwd()
				filepath = fmt.Sprintf("%s\\%s", path, strings.Trim(db, "./"))
			} else {
				filepath = db
			}

			_ = os.RemoveAll(filepath)
		}

		InitSqlite(c.String("db"), c.String("table"))
	case "memory":
		InitSqlite(":memory:", c.String("table"))
	case "database":

		table := c.String("table")

		InitSql(
			c.String("host"),
			c.String("port"),
			c.String("username"),
			c.String("password"),
			c.String("dbname"),
			table,
		)

		if force {
			if DB.Migrator().HasTable(table) {
				_ = DB.Migrator().DropTable(table)
			}
		}

		AutoMigrate()
	default:
		logrus.Error("无效的 driver")
		return errors.New("无效的 driver")
	}

	return nil
}
