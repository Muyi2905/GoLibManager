package models

import (
	"github.com/Muyi2905/GoLibManager/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `gorm : "column:name" json:name`
	Author      string `json:"author"`
	Publication string `json: "publication"`
}


func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}


func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}


func GetBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}


func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book 
	idData:= db.Where("Id=?", Id).Find(&getBook)
	return &getBook, idData
}

func DeleteByID(ID int64) *Book{
	var book Book
	 db.Where("Id=?", ID).Delete(book)
	return &book
}
