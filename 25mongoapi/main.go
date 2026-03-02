package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Ye7ia01/Study.Golang.MongoAPI/db"
	router "github.com/Ye7ia01/Study.Golang.MongoAPI/routers"
)

func main() {

	// 1 - Connect to MongoDB through My dbContext
	/* The Initialization can be in an "init" funciton to run once started */
	fmt.Println("Initiating DB Connection")
	dbClient, err := db.ConnectMongoDb()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB Successfully")
	defer dbClient.Disconnect(context.TODO())

	fmt.Println("Initializing Collection .. ")
	// 2 - Set the collection for the repository
	db.NewMovieCollection(dbClient.Database("netflix").Collection("mycollection"))
	fmt.Println("Initialized Repository COllection Instance")

	r := router.Router()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	/* Creates a goroutine automatically for each request */
	log.Fatal(srv.ListenAndServe())

	/*

		fmt.Println("Initializing Collection .. ")
		// 2 - Set the collection for the repository
		db.NewMovieCollection(dbClient.Database("netflix").Collection("mycollection"))

		fmt.Println("Adding Single Movie")
		// 3- Test the repository functions'
		// Create Single Movie
		movieId, err := db.CreateSingleMovie(&movie.Netflix{Name: "The Dark Knight", IsWatched: false})
		if err != nil {
			fmt.Println("Error creating movie : ", err)
		}
		fmt.Println("Inserted Movie ID : ", movieId)

		// Create Many Movies
		movies := []any{
			&movie.Netflix{Name: "Inception", IsWatched: false},
			&movie.Netflix{Name: "Interstellar", IsWatched: false},
		}

		fmt.Println("Adding Many Movies")
		movieIds, err := db.CreateManyMovies(movies)
		if err != nil {
			fmt.Println("Error creating movies : ", err)
		}

		fmt.Println("New Movies Created")

		fmt.Println("Updating Single Movie")
		// Update Single Movie
		updatedCount, err := db.UpdateMovie(movieId)
		if err != nil {
			fmt.Println("Error updating movie : ", err)
		}
		fmt.Printf("Updated %d Movie(s) \n", updatedCount)

		fmt.Println("Updating Many Movies")
		// Update Many Movies
		// ids := []string{movieId.(string)}
		updatedManyCount, err := db.UpdateManyMovies(movieIds)
		if err != nil {
			fmt.Println("Error updating movies : ", err)
		}
		fmt.Printf("Updated %d Movie(s) \n", updatedManyCount)

		fmt.Println("Deleting Single Movie")
		// Delete Movie
		err = db.DeleteMovie(movieId)
		if err != nil {
			fmt.Println("Error deleting movie : ", err)
		}
		fmt.Println("Movie Deleted")

		// Verify All Test works by Retrieving All Movies
		moviesList, err := db.GetAllMovies()
		if err != nil {
			fmt.Println("Error retrieving movies : ", err)
		}
		fmt.Println("Movies List : ", moviesList)

	*/

}
