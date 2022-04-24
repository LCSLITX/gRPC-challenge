package main

import (
	"os"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	mclient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("CLOUD_CONNECTION_STRING")))
	
	if err != nil {
		log.Fatalln(err)
	}
	
	err = mclient.Connect(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	
	MongoCollection = mclient.Database("cryptodb").Collection("crypto")

	log.Println("MongoDB ready.")
}
