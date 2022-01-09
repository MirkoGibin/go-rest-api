package main

import (
	"encoding/json"
	"go-rest-api/model"
	"go-rest-api/util"

	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (Model)
// book structure maps a json object

// Our mocked database. Collections are called "slice", which are more flexible arrays
var booksNumber = 10
var books []model.Book

// remember to init the module, go init go-rest-api
func main() {
	// init the mux router
	router := mux.NewRouter() // := is the operator for type inference

	// example of var declaration without type inference
	//	var specificRouter mux.Router = *mux.NewRouter()
	//specificRouter.HandleFunc("/api/books")
	intializeModel()
	// Router handlers and enpoints
	router.HandleFunc("/api/books/", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books/", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router)) // port setting and which router serves requests to this port. Handler wrapped in a log, in case of error this will be printed and the application will be stopped

}

func intializeModel() {
	for i := 0; i < booksNumber; i++ {
		var stringCounter string = strconv.Itoa(i)
		author := model.Author{Firstname: "name", Lastname: "surname"}
		books = append(books, model.Book{
			ID:     stringCounter,
			Isbn:   "isbn-0123456" + stringCounter,
			Title:  "Harry Potter " + stringCounter,
			Author: &author,
		})
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all books...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all books...")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result := model.Book{}
	for _, item := range books { //range is used to loop into a collection
		if item.ID == params["id"] {
			result = item
		}
	}

	json.NewEncoder(w).Encode(result)

}

func createBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new book...")
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book) // take the json body and map to the book object

	book.ID = util.NextId(books)
	books = append(books, book)

	log.Println(len(books))
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating book...")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]
	for index, item := range books {
		if item.ID == params["id"] {
			books[index] = book
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting book...")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) // i replace the content of the element i want to delete with the rest of the array. Everything is a reference
			break
		}
	}

	json.NewEncoder(w).Encode(books)

}
