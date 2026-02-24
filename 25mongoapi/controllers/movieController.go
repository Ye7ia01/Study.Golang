package movieController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ye7ia01/Study.Golang.MongoAPI/db"
	movie "github.com/Ye7ia01/Study.Golang.MongoAPI/models"
	"github.com/gorilla/mux"
)

func CreateMovieAction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movie *movie.Netflix
	// UnMarshal the Json Body
	fmt.Println("Decoding Json Movie")
	err := json.NewDecoder(r.Body).Decode(&movie)
	fmt.Println("Decode Movie ", movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println("Creating DB Movie")
	// Create the movie in the database
	newMovie, errCreate := db.CreateSingleMovie(movie)
	if errCreate != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errCreate)
		return
	}
	fmt.Println("Create Movie Successfully")
	// If everything goes well, send the new Body Json
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(newMovie)
}

func UpdateMovieDataAction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// Read Id for wha to be updated
	params := mux.Vars(r)
	fmt.Println("COntroller Received Id ", params["id"])

	// Read Movie objct to update
	var movie *movie.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	fmt.Println("Controller Decoded Movie as ", movie)

	fmt.Println("Starting Update in DB")
	// Update in Database
	updateCount, updateErr := db.UpdateMovieData(params["id"], movie)
	if updateErr != nil {
		fmt.Println("Return Error ", updateErr)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(updateErr)
		return
	}

	fmt.Println("Returned count ", updateCount)
	fmt.Println("Returned Error in Succes ", updateErr)

	// If every thing goes well
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(updateCount)
}

func UpdateMovieWatchedAction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// Get movie Id
	params := mux.Vars(r)

	// Update the movie in the DB
	count, err := db.UpdateMovie(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	// If every thing goes well , send number of items updated
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(count)
}

func DeleteMovieAction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// Try delete from the database
	err := db.DeleteMovie(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	// if every thing goes well
	w.WriteHeader(http.StatusAccepted)
}

func GetAllMovieAction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	movies, err := db.GetAllMovies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	// If every thing goes well
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(movies)
}
