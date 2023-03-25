package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "subham:subham235@tcp(127.0.0.1:3306)/bookstoreproject?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d

	fmt.Println("Connected to DB")
}

func GetDB() *gorm.DB {
	return db
}
