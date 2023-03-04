# go-city

全国「省市区县乡镇街道」数据来源于【京东】

# install

```shell
go install https://github.com/xiaoxuan6/go-city.git
```

# Usage

## Sqilte

```shell
go-city sync --db=./sqlite.db
```

## Database

```shell
go-city sync --host=127.0.0.1 --port=3306 --username=root --password=root --dbname=city
```

## More

```shell
NAME:
   go-city sync

USAGE:
   go-city sync [command options] [arguments...]

DESCRIPTION:
   同步省市区到数据库

OPTIONS:
   --driver value, -d value  db driver，Support：sqlite、database、memory (default: "sqlite") [%GO_CITY_DRIVER%]
   --db value                database file (default: "./sqlite.db") [%GO_CITY_DB%]
   --host value              database host (default: "127.0.0.1") [%GO_CITY_DB_HOST%]
   --port value              database port (default: "3306") [%GO_CITY_DB_PORT%]
   --username value          database username (default: "root") [%GO_CITY_DB_USERNAME%]
   --password value          database password (default: "root") [%GO_CITY_DB_PASSWORD%]
   --dbname value            database dbname [%GO_CITY_DB_NAME%]
   --help, -h                show help
```