package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book struct (model)
type book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *author `json:"author"`
}

type author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var b book
	json.NewDecoder(r.Body).Decode(&b)
	b.ID = strconv.Itoa(rand.Intn(9999))
	books = append(books, b)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			return
		}
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var b book
			json.NewDecoder(r.Body).Decode(&b)
			b.ID = params["id"]
			books = append(books, b)
			return
		}
	}

}

func main() {
	//Init router
	r := mux.NewRouter()

	// Mock data --@Todo DB
	books = append(books, book{ID: "1", Isbn: "4422", Title: "Book one", Author: &author{Firstname: "John", Lastname: "Abraham"}})
	books = append(books, book{ID: "2", Isbn: "6778", Title: "Book two", Author: &author{Firstname: "Ricky", Lastname: "Pointing"}})

	//Route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
