package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID        string    `json:"id"`
	Genre     string    `json:"genre"`
	Title     string    `json:"title"`
	Directors *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			var updatedMovie Movie
			if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatedMovie.ID = params["id"]
			movies[index] = updatedMovie
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Genre: "Drama", Title: "The Godfather", Directors: &Director{FirstName: "Francis", LastName: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "2", Genre: "Thriller", Title: "Psycho", Directors: &Director{FirstName: "Alfred", LastName: "Hitchcock"}})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the Movies API!"})
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
