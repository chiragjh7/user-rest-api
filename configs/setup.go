package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB connects to the database
func ConnectDB() *mongo.Client {

	// Set client options
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}
	
	// set the context for the request
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	// defer the cancel() to avoid memory leaks
	defer cancel()

	// return the client
	return client
}

// DB is the database connection
var DB *mongo.Client = ConnectDB()

// GetCollection returns the collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("user-api").Collection(collectionName)
	return collection
}
