package main

import(
	"os"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getEnv(key, standard string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = standard
	}
	return value
}

func init() {
	mclient, err := mongo.NewClient(options.Client().ApplyURI(getEnv("CLOUD_CONNECTION_STRING", "mongodb://root:root@localhost:27017/")))
	
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
