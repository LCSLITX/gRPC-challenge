package main

import (
	"fmt"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/lucassauro/klever-challenge/proto"
)

// UpdateCrypto function receives a context and request with Crypto struct as parameters, update the coin in the database and returns its id and error.
func (s *Server) UpdateCrypto(ctx context.Context, req *pb.Crypto) (*pb.CryptoId, error) {

	erro := ValidateCoin(req)

	if erro != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln(erro))
	}

	id := bson.M{"_id": req.Id}

	updateTo := bson.M{
		"name": req.Name,
		"short": req.Short,
		"votes": req.Votes,
	}

	res := MongoCollection.FindOneAndUpdate(
		ctx, 
		id, bson.M{"$set": updateTo}, 
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	coin := &Coin{}

	if err := res.Decode(coin); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("No coin with specified Id. %v", err))
	}

	return &pb.CryptoId{
		Id: coin.Id,
	}, nil
}