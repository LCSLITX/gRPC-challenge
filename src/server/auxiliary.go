package main

import (
	"fmt"
	"log"
	"errors"
	"context"

	"github.com/subosito/gotenv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-playground/validator/v10"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

// MongoCollection incialization.
var MongoCollection *mongo.Collection

// Validate inicialization.
var Validate *validator.Validate

// The following function doesn't allowed Go compiler to build a binary file in an environment without a .env file.
// So, the solution was simply comment the code that tries to access .env variables.
// gotenv inicialization.
func init() {
	if err := gotenv.Load(); err != nil {
		log.Fatalln(err)
	}
}

// ===== Structs =====.
type Coin struct {
	ID 					uint32			`bson:"_id,omitempty" json:"id,omitempty" validate:"numeric"`
	Name 				string 			`bson:"name,omitempty" json:"name,omitempty" validate:"required,alpha"`
	Short 			string 			`bson:"short,omitempty" json:"short,omitempty" validate:"required,alphanum"`
	Votes 			int64 			`bson:"votes,omitempty" json:"votes,omitempty" validate:"numeric"`
}

type Votes struct {
	Votes 			int64 			`bson:"votes" json:"votes"`
}

// ===== Helper functions =====.

// FindLastId function finds the Id of the last added document in a mongo collection.
func FindLastID() (uint32, error) {
	opts := options.FindOne()

	opts.SetSort( bson.D { { Key: "_id", Value: -1 } } )

	filter := bson.D{}

	lastID := MongoCollection.FindOne(context.Background(), filter, opts)

	data := &Coin{}

	if err := lastID.Decode(data); errors.Is(err, mongo.ErrNoDocuments) {
		return 0, nil
	} else if err != nil {
		return 0, status.Errorf(codes.Internal, fmt.Sprintln(err))
	}

	return data.ID, nil
}

// FrankensteinGetVotesPls function receives a parameter of type bson.M, returned by mongo.Watch(),
// and iterates over it until it get number of votes to return it.
// know more: https://stackoverflow.com/questions/61354850/how-to-assert-a-primitive-m-to-a-mapstringstring.
func FrankensteinGetVotesPls(data bson.M) int64 {
	mapa := make(map[string]interface{})
	
	for key, value := range data["updateDescription"].(primitive.M) {
		if key == "updatedFields" {
			mapa["a"] = value
		}
	}
	
	for key, value := range mapa["a"].(primitive.M) {
		if key == "votes" {
			mapa["a"] = value
		}
	}

	// type assertion.
	// t := mapa["a"].(int64)
	t := mapa["a"].(int32)
	// interface conversion problem. Know more: 
	// https://stackoverflow.com/questions/70705673/panic-interface-conversion-interface-is-float64-not-int64

	return int64(t)
}

// DoesThisCryptoExist function receives parameter of type *pb.CryptoId,
// search mongo collection for it and return info Coin struct.
func DoesThisCryptoExist(id uint32) (*Coin, error) {
	filter := bson.M{"_id": id}
	
	res := MongoCollection.FindOne(context.Background(), filter)
	
	if res.Err() != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln(res.Err()))
	}

	coin := &Coin{}

	if err := res.Decode(coin); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln(err))
	}

	return coin, nil
}

// ValidateCoin function receives a pb.Crypto as parameter and executes validator
// to check whether the coin is valid or not. Return error only. 
func ValidateCoin(coin *pb.Crypto) error {
	isCoinValid := &Coin {
		ID: coin.Id,
		Name: coin.Name,
		Short: coin.Short,
		Votes: coin.Votes,
	}

	// validate existence of both name and short 
	Validate = validator.New()

	validationErr := Validate.Struct(isCoinValid)
	
	if validationErr != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintln(validationErr))
	}

	return nil
}



