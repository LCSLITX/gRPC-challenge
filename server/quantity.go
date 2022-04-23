package main

import (
	"log"
	"context"
	
	"go.mongodb.org/mongo-driver/bson"
)

func QuantityOfCryptos() int64 {
	count, err := MongoCollection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return count
}