package models

import (
	"github.com/cokikip/go-bookstore/pkg/configs"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Auther      string `json:"auther"`
	Publication string `json:"publication"`
}

func init() {
	configs.Connect()
	db = configs.GetDb()
	db.AutoMigrate(&Book{})
}
