package routes

import (
	"github.com/gorilla/mux"
	  "github.com/Muyi2905/GoLibManager/pkg/controllers"
)

var RegisterBookstoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateById).Methods("PATCH")
	router.HandleFunc("/book/{id}", controllers.DeleteByID).Methods("DELETE")
}

