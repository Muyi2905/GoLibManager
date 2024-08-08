package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Muyi2905/GoLibManager/pkg/config"
)

type Book struct {
	gorm.Model

	Name string `gorm : "" json:name`
	Author string `json:"author"`
	Publication string `json: "publication"`

}