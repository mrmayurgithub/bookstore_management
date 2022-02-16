package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root@localhost:password@/simplerest?charset=utf8&parseTime=True&loc=LOCAL")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
