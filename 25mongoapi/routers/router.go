package router

import (
	movieController "github.com/Ye7ia01/Study.Golang.MongoAPI/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()

	// Create
	r.HandleFunc("/api/movies", movieController.CreateMovieAction).Methods("POST")
	// Get All
	r.HandleFunc("/api/movies", movieController.GetAllMovieAction).Methods("GET")
	// Update Watched
	r.HandleFunc("/api/movies/watched/{id}", movieController.UpdateMovieWatchedAction).Methods("PUT")
	// Update Data
	r.HandleFunc("/api/movies/{id}", movieController.UpdateMovieDataAction).Methods("PUT")
	// Delte
	r.HandleFunc("/api/movies/{id}", movieController.DeleteMovieAction).Methods("DELETE")

	return r
}
