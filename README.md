# go-city

全国「省市区县乡镇街道」，数据来源于【京东】

# install

```shell
go install github.com/xiaoxuan6/go-city@latest
```

# Usage

## Sqlite

```shell
go-city sync --db=./sqlite.db
```

## Database

```shell
go-city sync --driver=database --host=127.0.0.1 --port=3306 --username=root --password=root --dbname=city --table=city
```

## More

```shell
NAME:
   go-city - go city

USAGE:
   go-city [global options] command [command options] [arguments...]

COMMANDS:
   sync, s     sync data to database, data source is jd
   version, v  show go city version
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
  --help, -h  show help

OPTIONS:
   --driver value, -d value  db driver，Support：sqlite、database、memory (default: "sqlite") [%GO_CITY_DRIVER%]
   --db value                database file (default: "./sqlite.db") [%GO_CITY_DB%]
   --host value              database host (default: "127.0.0.1") [%GO_CITY_DB_HOST%]
   --port value              database port (default: "3306") [%GO_CITY_DB_PORT%]
   --username value          database username (default: "root") [%GO_CITY_DB_USERNAME%]
   --password value          database password (default: "root") [%GO_CITY_DB_PASSWORD%]
   --dbname value            database dbname [%GO_CITY_DB_NAME%]
   --table value             database table name (default: "cities") [%GO_CITY_DB_TABLE_NAME%]
   --force                   drop table (default: false) [%GO_CITY_DB_FORCE%]
   --help, -h                show help
```
