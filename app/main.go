package main

import (
	"app/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	router := mux.NewRouter()

	// Route Handlers/Endpoint
	router.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/api/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controller.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
