package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Muyi2905/GoLibManager/pkg/models"
	"github.com/Muyi2905/GoLibManager/pkg/utils"
	"github.com/gorilla/mux"
)

var Newbook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBook()
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "error marshling json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error getting book id", http.StatusNotFound)
	}
	bookdetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := models.Book{}
	utils.ParseForm(r, &createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error deleting bookId", http.StatusNotFound)
	}
	book := models.DeleteByID(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	updateBook := models.Book{}
	vars := mux.Vars(r)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("error updating book", http.StatusNotFound)
	}
	bookdeatils, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookdeatils.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookdeatils.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookdeatils.Publication = updateBook.Publication
	}

	db.Save(&bookdeatils)
	res, _ := json.Marshal(bookdeatils)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
