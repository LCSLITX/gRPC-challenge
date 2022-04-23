package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/lucassauro/klever-challenge/proto"
)

func (s *Server) ReadCrypto(ctx context.Context, req *pb.CryptoId) (*pb.Crypto, error) {
	id := bson.M{"_id": req.Id}

	res := MongoCollection.FindOne(ctx, id)

	coin := &Coin{}
	if err := res.Decode(coin); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("No coin with specified Id. %v", err))
	}

	return &pb.Crypto{
		Id: coin.Id,
		Name: coin.Name,
		Short: coin.Short,
		Votes: coin.Votes,
	}, nil
}
