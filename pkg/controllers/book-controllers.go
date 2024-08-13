package controllers

import (
	"encoding/json"
	"fmt"
	"log"
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
		fmt.Println("Error getting book id", httpstatus)
	}
	bookdetails,_ := models.GetBookById(ID)
	res,_ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
