package main

import (
	"fmt"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"

	pb "github.com/lucassauro/klever-challenge/proto"
)

// DownvoteCrypto receives context and request with crypto id as parameters. Returns CryptoVotes struct containing the updated number of votes.
func (s *Server) DownvoteCrypto(ctx context.Context, req *pb.CryptoId) (*pb.CryptoVotes, error) {
	id := bson.M{"_id": req.Id}

	res := MongoCollection.FindOneAndUpdate(
		ctx, 
		id, 
		bson.D { primitive.E { Key: "$inc", Value: bson.D { primitive.E { Key: "votes", Value: -1 } } } }, 
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	coin := &Coin{}

	if err := res.Decode(coin); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln(err))
	}

	return &pb.CryptoVotes{
		Votes: coin.Votes,
	}, nil
}