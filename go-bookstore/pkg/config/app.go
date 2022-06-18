package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Pruebas2022@(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error while connecting to mysql")
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
