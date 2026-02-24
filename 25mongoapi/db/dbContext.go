package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName string = "netflix"
const dbColl string = "mycollection"

var collection mongo.Collection

func ConnectMongoDb() (*mongo.Client, error) {

	uri := os.Getenv("go_mongo_db_uri")

	clientOptions := options.Client().ApplyURI(uri)
	//.setBsonOptions(bsonOptions)
	// extra bson options can be used for configuring marshaling and
	// unmarshaling of bson

	/* Context is a connection context parameters like timeout, close , background */
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		// Log Detailed error and Exit
		return nil, err
	}

	return client, nil
}

// collection = *client.Database(dbName).Collection(dbColl)
// fmt.Println("Connection to mongo DB is successfull")
//
// }
