package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client

func GetCollection(collectionName string) *mongo.Collection {
	c := GetConfig()
	collection := client.Database(c.DBName).Collection(collectionName)
	return collection
}

func ConnectDB() (err error) {
	uri := GetURI()
	cl, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	ctx := context.Background()
	err = cl.Connect(ctx)
	if err != nil {
		return
	}

	err = cl.Ping(ctx, nil)
	if err != nil {
		return
	}
	client = cl
	log.Println("Connected to MongoDB")
	return nil
}
