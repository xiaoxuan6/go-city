# go-city

全国「省市区县乡镇街道」数据来源于【京东】

# install

```shell
go install https://github.com/xiaoxuan6/go-city.git
```

# Usage

```shell
NAME:
   go-city sync

USAGE:
   go-city sync [command options] [arguments...]

DESCRIPTION:
   同步省市区到数据库

OPTIONS:
   --driver value, -d value  db driver，Support：sqlite、database、memory (default: "sqlite") [%GO_CITY_DRIVER%]
   --db value                database file (default: "/sqlite.db") [%GO_CITY_DB%]
   --help, -h                show help
```