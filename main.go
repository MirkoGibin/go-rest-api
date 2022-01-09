package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (Model)
// book structure maps a json object
type Book struct {
	ID     string  `json:"id"` // ID field maps the field id of the json object
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Our mocked database. Collections are called "slice", which are more flexible arrays
var booksNumber = 10
var books []Book

// remember to init the module, go init go-rest-api
func main() {
	// init the mux router
	router := mux.NewRouter() // := is the operator for type inference

	// example of var declaration without type inference
	//	var specificRouter mux.Router = *mux.NewRouter()
	//specificRouter.HandleFunc("/api/books")
	intializeModel()
	// Router handlers and enpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router)) // port setting and which router serves requests to this port. Handler wrapped in a log, in case of error this will be printed and the application will be stopped

}

func intializeModel() {
	for i := 0; i < booksNumber; i++ {
		var stringCounter string = strconv.Itoa(i)
		books = append(books, Book{
			ID:     stringCounter,
			Isbn:   "isbn-0123456" + stringCounter,
			Title:  "Harry Potter " + stringCounter,
			Author: &Author{Firstname: "name", Lastname: "surname"},
		})
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all books...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}
