package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// func QuantityOfCryptos() int64 {
// 	count, err := MongoCollection.CountDocuments(context.Background(), bson.D{})
// 	if err != nil {
// 		log.Fatalf("Error: %v", err)
// 	}
// 	return count
// }

func FindLastId() int64 {
	opts := options.FindOne()
	opts.SetSort(bson.D{{Key: "_id", Value: -1}})
	filter := bson.D{}
	lastId := MongoCollection.FindOne(context.Background(), filter, opts)
	data := &Coin{}
	if err := lastId.Decode(data); err != nil {
		status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}
	return int64(data.Id)
}
