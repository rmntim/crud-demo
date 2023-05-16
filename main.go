package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(movies) > 0 {
		json.NewEncoder(w).Encode(movies)
	} else {
		fmt.Fprint(w, "No movies found :(")
	}
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	fmt.Fprintf(w, "Movie id `%s` does not exist :(", params["id"])
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/`json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		fmt.Fprint(w, "Bad format!")
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/`json")
	params := mux.Vars(r)
	var newMovie Movie
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		fmt.Fprint(w, "Bad format!")
		return
	}
	newMovie.ID = params["id"]
	for idx, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(append(movies[:idx], newMovie), movies[idx+1:]...)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	fmt.Fprintf(w, "Movie id `%s` does not exist :(\n", params["id"])
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	fmt.Fprintf(w, "Movie id `%s` does not exist :(\n", params["id"])
	json.NewEncoder(w).Encode(movies)
}

func main() {
	movies = append(
		movies,
		Movie{
			ID:       "1",
			Isbn:     "123123",
			Title:    "Doctor Strange",
			Director: &Director{Firstname: "Zach", Lastname: "Snyder"},
		},
		Movie{
			ID:       "2",
			Isbn:     "456456",
			Title:    "Spider Man",
			Director: &Director{Firstname: "Alan", Lastname: "Wake"},
		},
	)

	r := mux.NewRouter()

	r.HandleFunc("/", indexPage).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	port := 8000
	fmt.Printf("Starting server at port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
