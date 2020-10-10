package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) // Get params

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}

	json.NewEncoder(writer).Encode(&Book{})
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID - not safe
			books = append(books, book)
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
	json.NewEncoder(writer).Encode(books)
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(books)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Data - @todo implement DB
	books = append(
		books,
		Book{
			ID:    "1",
			Isbn:  "543563",
			Title: "Percy Jackson",
			Author: &Author{
				Firstname: "John",
				Lastname:  "Snow",
			},
		},
	)

	books = append(
		books,
		Book{
			ID:    "2",
			Isbn:  "323254",
			Title: "King Arthur",
			Author: &Author{
				Firstname: "Merlin",
				Lastname:  "Lady's Lake",
			},
		},
	)

	// Route Handlers/Endpoint
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
