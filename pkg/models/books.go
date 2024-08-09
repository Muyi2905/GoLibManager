package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Muyi2905/GoLibManager/pkg/config"
)

type Book struct {
	gorm.Model

	Name string `gorm : "column:name" json:name`
	Author string `json:"author"`
	Publication string `json: "publication"`

}

func init (){
	config.Connect()
	db:= config.GetDB()
	db.AutoMigrate(&Book{})
}