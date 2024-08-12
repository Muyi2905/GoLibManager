package main

import (
	"log"
	"net/http"

	"github.com/Muyi2905/GoLibManager/pkg/routes"
	"github.com/github.com/Muyi2905/GoLibManager/pkg/routes"
	"github.com/gorilla/mux"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
http.Handle("/", r)
log.Fatal(http.ListenAndServe(":8080", r))	
}