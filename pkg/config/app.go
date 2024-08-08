package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func connect(){
	d, err := gorm.Open("mysql", "muyiwa:Muyi2905@/golang?charset=utf8&parseTime=True&loc=Local")
		if err!= nil{
			panic(err)
		}
db = d	
}

func GetDB () *gorm.DB{
	return db
}