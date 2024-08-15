package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbParseTime := os.Getenv("DB_PARSE_TIME")
	dbLoc := os.Getenv("DB_LOC")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbParseTime, dbLoc)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
