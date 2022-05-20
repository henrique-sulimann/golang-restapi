package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartMongo() *mongo.Client {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASS") + "@" + os.Getenv("DATABASE_SERVER") + ":" + os.Getenv("DATABASE_PORT")))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:123456@localhost:27023"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var MONGO *mongo.Client = StartMongo()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("games").Collection(collectionName)
	return collection
}
