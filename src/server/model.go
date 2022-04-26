package main

import(
	"os"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// The following function doesn't allowed Go compiler to build a binary file in an environment without a .env file.
// So, the solution was simply comment the code that tries to access .env variables.
func getEnv(key, standard string) string {
	value, exists := os.LookupEnv(key)

	if _, ok := os.LookupEnv("TEST"); ok {
		log.Println("Using local database.")

		value = standard
	}

	if !exists {
		value = standard
	}

	return value
}


func init() {
	connectedTo := getEnv("CLOUD_CONNECTION","mongodb://root:root@localhost:27017/")
	// The original code uses .env to access a database hosted on https://cloud.mongodb.com.
	// Unfortunately, Go compiler doesn't build a binary file if there is no .env file to be accessed.
	mclient, err := mongo.NewClient(options.Client().ApplyURI(connectedTo))
	// mclient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	
	if err != nil {
		log.Fatalln(err)
	}
	
	err = mclient.Connect(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	
	MongoCollection = mclient.Database("cryptodb").Collection("crypto")

	log.Printf("MongoDB ready. Connected to: %v", connectedTo)
}
