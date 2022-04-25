package main

import (
	"fmt"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	pb "github.com/lucassauro/klever-challenge/proto"
)

// LiveCryptoVotes receives a request with id and a stream as parameters; and it streams every update on votes of the specific crypto.
func (s *Server) LiveCryptoVotes(req *pb.CryptoId, stream pb.CryptoService_LiveCryptoVotesServer) error {

	match := bson.D { primitive.E { Key: "$match", Value: bson.D { primitive.E { Key: "documentKey", Value: bson.D { primitive.E { Key: "_id", Value: req.Id } } } } } } 
	
	watch, err := MongoCollection.Watch(context.TODO(), mongo.Pipeline{match})
	
	if err != nil {
		status.Errorf(codes.NotFound, fmt.Sprintln(err))
	}
	
	defer watch.Close(context.Background())	
	
	for watch.Next(context.Background()){
		var data bson.M

		if err := watch.Decode(&data); err != nil {
			status.Errorf(codes.Internal, fmt.Sprintln(err))	
		}
		
		votesNumber := FrankensteinGetVotesPls(data)
		
		stream.Send(&pb.CryptoVotes{
			Votes: votesNumber,
		})
	}
	
	return nil
}
