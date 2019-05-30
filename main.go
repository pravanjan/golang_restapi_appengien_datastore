package main

import (
	"fmt"
	"golang_restapi_appengien_datastore/controllers"
	"net/http"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func main() {
	fmt.Println("inside main method ")

	router := mux.NewRouter()
	router.HandleFunc("/api/entry/create", controllers.CreateNewEntry).Methods("POST")
	router.HandleFunc("/", controllers.DefaultView).Methods("GET")

	
	http.Handle("/", router)
	appengine.Main()

}
