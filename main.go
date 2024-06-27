package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
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

// @title Movies API
// @version 1.0
// @description This is a simple API for managing a collection of movies.
// @contact.name API Support
// @contact.email support@example.com

// getMovies godoc
// @summary Get all movies
// @description Get details of all movies
// @tags movies
// @produce json
// @success 200 {array} Movie
// @router /movies [get]
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// getMovie godoc
// @summary Get a movie by ID
// @description Get details of a movie by ID
// @tags movies
// @produce json
// @param id path string true "Movie ID"
// @success 200 {object} Movie
// @failure 404 {object} string "Movie not found"
// @router /movies/{id} [get]
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

// createMovie godoc
// @summary Create a new movie
// @description Create a new movie with a random ID
// @tags movies
// @accept json
// @produce json
// @param movie body Movie true "New movie details"
// @success 201 {object} Movie
// @failure 400 {object} string "Bad request"
// @router /movies [post]
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

// updateMovie godoc
// @summary Update a movie by ID
// @description Update details of a movie by ID
// @tags movies
// @accept json
// @produce json
// @param id path string true "Movie ID"
// @param movie body Movie true "Updated movie details"
// @success 200 {object} Movie
// @failure 400 {object} string "Bad request"
// @failure 404 {object} string "Movie not found"
// @router /movies/{id} [put]
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

// deleteMovie godoc
// @summary Delete a movie by ID
// @description Delete a movie by ID
// @tags movies
// @produce json
// @param id path string true "Movie ID"
// @success 200 {array} Movie
// @failure 404 {object} string "Movie not found"
// @router /movies/{id} [delete]
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

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the Movies API!"})
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)

	// Serve the Swagger spec
	r.Path("/swagger.yaml").Handler(http.FileServer(http.Dir("./")))

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
