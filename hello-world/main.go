// Playing areound with APIs

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie item
type Movie struct {
	Title   string `json:"title"`
	Year    int    `json:"pub_year"`
	Content string `json:"content"`
}

// Movies list
type Movies []Movie

func allMovies(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{Title: "Alien", Year: 1979, Content: "Ellen Ripley muore"},
	}

	fmt.Println("All movies")
	json.NewEncoder(w).Encode(movies)
}

func postMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post movies work")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handleRequests() {

	MyRouter := mux.NewRouter().StrictSlash(true)
	MyRouter.HandleFunc("/", homepage)
	MyRouter.HandleFunc("/movies", allMovies).Methods("GET")
	MyRouter.HandleFunc("/movies", postMovies).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", MyRouter))
}

func main() {
	handleRequests()
}
