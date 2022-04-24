package main

import (
	"fmt"
	"context"

	"github.com/subosito/gotenv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo/options"	

)

// global MongoCollection incialization
var MongoCollection *mongo.Collection

// global validator inicialization
var validate *validator.Validate

// global gotenv inicialization.
func init() {
	gotenv.Load()
}

// ===== Structs =====
type Coin struct {
	Id 					uint32			`bson:"_id,omitempty" json:"id,omitempty"`
	Name 				string 			`bson:"name,omitempty" json:"name,omitempty" validate:"required"`
	Short 			string 			`bson:"short,omitempty" json:"short,omitempty" validate:"required"`
	Votes 			int64 			`bson:"votes,omitempty" json:"votes,omitempty"`
}

type Votes struct {
	Votes 			int64 			`bson:"votes" json:"votes"`
}

// ===== Helper functions =====

// FindLastId function finds the Id of the last added document in a mongo collection.
func FindLastId() (uint32, error) {
	opts := options.FindOne()
	opts.SetSort(bson.D{{Key: "_id", Value: -1}})
	filter := bson.D{}
	lastId := MongoCollection.FindOne(context.Background(), filter, opts)
	data := &Coin{}

	if err := lastId.Decode(data); err != nil {
		return 0, status.Errorf(codes.Internal, fmt.Sprintln(err))
	}
	return uint32(data.Id), nil
}

// FrankensteinGetVotesPls function receives a parameter of type bson.M, returned by mongo.Watch(), and iterates over it until it get number of votes to return it.
// know more: https://stackoverflow.com/questions/61354850/how-to-assert-a-primitive-m-to-a-mapstringstring
func FrankensteinGetVotesPls(data bson.M) int64 {
	a := make(map[string]interface{})
	for key, value := range data["updateDescription"].(primitive.M) {
		if key == "updatedFields" {
			a["1"] = value
		}
	}
	for key, value := range a["1"].(primitive.M) {
		if key == "votes" {
			a["1"] = value
		}
	}
	// type assertion
	t := a["1"].(int64)

	return t
}

// DoesThisCryptoExist function receives parameter of type *pb.CryptoId, search mongo collection for it and return info Coin struct.
func DoesThisCryptoExist(id uint32) (*Coin, error) {
	
	filter := bson.M{"_id": id}
	
	res:= MongoCollection.FindOne(context.Background(), filter)
	
	if res.Err() != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln(res.Err()))
	}

	coin := &Coin{}

	if err := res.Decode(coin); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln(err))
	}

	return coin, nil
}