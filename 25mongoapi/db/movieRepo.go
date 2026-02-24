package db

import (
	"context"
	"fmt"

	movie "github.com/Ye7ia01/Study.Golang.MongoAPI/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var movieCollection mongo.Collection

func NewMovieCollection(collection *mongo.Collection) {
	if collection != nil {
		movieCollection = *collection
	}
}

func CreateSingleMovie(movie *movie.Netflix) (interface{}, error) {
	inserted, err := movieCollection.InsertOne(context.TODO(), movie)
	// Extra options can be passed for the InsertOne operation

	if err != nil {
		return nil, err
	}

	return inserted.InsertedID, nil
}

func CreateManyMovies(movies []any) ([]interface{}, error) {

	// Make sure the array is of of type *movie.Netflix
	for _, value := range movies {
		if _, ok := value.(*movie.Netflix); !ok {
			return nil, fmt.Errorf("Invalid Type")
		}
	}

	// InsertMany for all received records
	inserted, err := movieCollection.InsertMany(context.TODO(), movies)
	// Extra options can be passed for the InsertMany operation

	if err != nil {
		// Make sure to read inserted Ids,
		// Since some exceptions happen and some records are already insterted
		return inserted.InsertedIDs, err
	}

	return inserted.InsertedIDs, nil
}

func UpdateMovie(_id string) (int64, error) {

	/*
		Any Update op expects 3 parameters :
			Query Filter (Where (docs/rows) the change will take place) (WHERE in SQL)
			Change Document (What key/columns will be changed)
			Options
	*/

	/*
			Update operator always in this format :
			bson.D{{"<update operator>", bson.D{{"<field>", <value>},
		                                    {"<field>", <value>}, ... }},
		       {"<update operator>", ... }, ... }
	*/
	// where the change will happen
	movieId, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return 0, err
	}
	filter := bson.D{{Key: "_id", Value: movieId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "is_watched", Value: true}}}}
	// Extra options can be passed for the updateOne operation
	// Upsert is to Insert If not found (in this case is just a demo)
	/*
		If the filter had more than one record
		The updateOne will update the first record retrieved,
		 Sorting may be applied in options to update the desired record only
	*/
	opts := options.Update().SetUpsert(true)
	mod, err := movieCollection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		return mod.ModifiedCount, err
	}

	return mod.ModifiedCount, nil
}

// Unlike the Original update which updates watching only
// This will update the data objct
func UpdateMovieData(id string, movie *movie.Netflix) (int64, error) {

	// Type Case Object ID
	_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, err
	}

	filter := bson.D{{Key: "_id", Value: _id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "movieName", Value: movie.Name},
		{Key: "isWatched", Value: movie.IsWatched},
		{Key: "starName", Value: movie.Star.Name},
		{Key: "starAge", Value: movie.Star.Age},
	}}}

	// opts := options.Update().SetUpsert(true)
	updateRes, updateErr := movieCollection.UpdateOne(context.TODO(), filter, update)
	if updateErr != nil {
		return 0, updateErr
	}
	// Successfully Modify
	return updateRes.ModifiedCount, nil
}

func UpdateManyMovies(_ids []string) (int64, error) {

	var movieIds []primitive.ObjectID
	for _, value := range _ids {
		movieId, err := primitive.ObjectIDFromHex(value)
		if err != nil {
			return 0, err
		}
		movieIds = append(movieIds, movieId)
	}

	// Where the change will happen
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: movieIds}}}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "is_watched", Value: true}}}}
	opts := options.Update().SetUpsert(true)

	updateRes, err := movieCollection.UpdateMany(context.TODO(), filter, update, opts)
	if err != nil {
		return updateRes.ModifiedCount, err
	}

	return updateRes.ModifiedCount, nil

}

func DeleteMovie(_id string) error {

	movieId, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return err
	}

	_, delErr := movieCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: movieId}})
	if delErr != nil {
		return err
	}

	return nil
}

func DeleteAllMovies() (int64, error) {
	delResult, delErr := movieCollection.DeleteMany(context.TODO(), bson.D{{}})
	if delErr != nil {
		/* In case some of the records are deleted before error */
		return delResult.DeletedCount, delErr
	}

	return delResult.DeletedCount, nil
}

func GetAllMovies() ([]movie.Netflix, error) {

	// set sort option to be ascending by name
	opts := options.Find().SetSort(bson.D{{Key: "name", Value: 1}})
	cursor, err := movieCollection.Find(context.TODO(), bson.D{{}}, opts)
	if err != nil {
		return nil, err
	}

	var movies []movie.Netflix
	if err = cursor.All(context.TODO(), &movies); err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	return movies, nil
}

func GetMovie(_id string) (interface{}, error) {

	movieId, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: movieId}}

	var movie *movie.Netflix
	findErr := movieCollection.FindOne(context.TODO(), filter).Decode(&movie)
	if findErr != nil {
		return nil, err
	}

	return movie, err

}
