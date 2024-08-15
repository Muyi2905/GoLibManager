package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql" // Required for MySQL dialect
    "log"
)

var db *gorm.DB

func Connect() {
    var err error

    d, err := gorm.Open("mysql", "muyiwa:Muyi2905@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    db = d
}

func GetDB() *gorm.DB {
    return db
}
