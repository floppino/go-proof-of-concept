package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// repo that let us differentiate APIs calls POST,GET etc
	"github.com/gorilla/mux"
)

// Movie item, creating a struct of type Movie with the following attributes
type Movie struct {
	Title   string `json:"title"`
	Year    int    `json:"pub_year"`
	Content string `json:"content"`
}

// Movies list, creating a list that will contain the Movies
type Movies []Movie

// Lists all movies in the body of the function
func allMovies(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{Title: "Alien", Year: 1979, Content: "Ellen Ripley muore"},
	}

	fmt.Println("All movies")
	// Encode: convert data to and from byte-level and textual representations
	json.NewEncoder(w).Encode(movies)
}

// Post call to insert our json via postman
func postMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post call works... ?")
}

// Homepage for our root endpoint
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Github!")
}

func handleRequests() {
	// MyRouter: let us use different APIs on the same url (GET,POST etc)
	MyRouter := mux.NewRouter().StrictSlash(true)
	// Routes
	MyRouter.HandleFunc("/", homepage)
	MyRouter.HandleFunc("/movies", allMovies).Methods("GET")
	MyRouter.HandleFunc("/movies", postMovies).Methods("POST")
	// Open the server on that port, log.Fatal listens for errors and closes the connection if needed
	log.Fatal(http.ListenAndServe(":8081", MyRouter))
}

func main() {
	handleRequests()
}
