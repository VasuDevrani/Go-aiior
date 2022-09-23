package main

import (
	"fmt"
	"log" //for formatted outputs like errors
	"math/rand"
	"net/http" //for setting up the htttp server
	"strconv"

	"github.com/gorilla/mux" //for route req, res mechanism similr to express

	// "strconv"
	// "math/rand"
	"encoding/json"
)

// models in the form of structs
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `jaon:"lastname"`
}

// database array
var movies []Movie

// controllers
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Cotent-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func updateMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>This is the home page</h1>"))
}

func main(){
	// creator the router instance
	r := mux.NewRouter();

	movies = append(movies, Movie{ID: "1", Isbn: "213324", Title: "First Movie", Director: &Director{Firstname: "Roman", Lastname: "Range"}})

	movies = append(movies, Movie{ID: "2", Isbn: "213876", Title: "Second Movie", Director: &Director{Firstname: "rahul", Lastname: "bishnoi"}})

	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at 8080 \n");
	log.Fatal(http.ListenAndServe(":8080", r));
}
