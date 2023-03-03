package command

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitSqlite(dsn string) {
	DB, _ = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	_ = DB.AutoMigrate(Response{})
}

func InitSql(host, port, username, password, dbname string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("数据库链接失败：%s", err.Error()))
	}

	DB = db
	_ = DB.AutoMigrate(Response{})
}

type Response struct {
	ID   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"varchar(125);not null;comment:'城市名称'"`
	Pid  int    `json:"pid" gorm:"int(10);not null;comment:'父级id'"`
}

func (r Response) TableName() string {
	return "cities"
}

func save(response ...Response) {
	if err := DB.Model(&Response{}).Create(&response).Error; err != nil {
		logrus.Error(fmt.Sprintf("插入失败：%s", err.Error()))
	}
}
