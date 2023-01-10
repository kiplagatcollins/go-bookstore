package configs

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	password := os.Getenv("PASSWORD")
	username := os.Getenv("USERNAME")
	mysqlConnection := fmt.Sprintf("%s:%s/bookstore?charset=utf8&parseTime=True&loc=Local", username, password)
	d, err := gorm.Open("mysql", mysqlConnection)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db

}
