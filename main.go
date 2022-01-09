package main

import (
	"log"
	"net/http"

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

// remember to init the module, go init go-rest-api
func main() {
	// init the mux router
	router := mux.NewRouter() // := is the operator for type inference

	// example of var declaration without type inference
	//	var specificRouter mux.Router = *mux.NewRouter()
	//specificRouter.HandleFunc("/api/books")

	// Router handlers and enpoints

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router)) // port setting and which router serves requests to this port. Handler wrapped in a log, in case of error this will be printed and the application will be stopped

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all books...")
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}
