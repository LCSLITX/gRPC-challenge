package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	log.Println("MongoDB ready.")
	mclient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = mclient.Connect(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	MongoCollection = mclient.Database("cryptodb").Collection("crypto")
}
