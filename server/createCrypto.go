package main

import (
	"fmt"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/lucassauro/klever-challenge/proto"
)

// CreateCrypto receives context and request with NewCrypto struct as parameters, create a coin in the database and return its auto incremented integer Id.
func (s *Server) CreateCrypto(ctx context.Context, req *pb.NewCrypto) (*pb.CryptoId, error) {
	// Following variable stores the Id of the last inserted document plus one.
	coinId := uint32(FindLastId() + 1)
	
	newCoin := &Coin {
		Id: coinId,
		Name: req.Name,
		Short: req.Short,
		Votes: 0,
	}
	
	_, err := MongoCollection.InsertOne(ctx, newCoin)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	return &pb.CryptoId{
		Id: coinId,
	}, nil
}
