package main

import (
	"fmt"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/lucassauro/klever-challenge/proto"
)

func (s *Server) CreateCrypto(ctx context.Context, req *pb.NewCrypto) (*pb.CryptoId, error) {	
	quant := uint32(FindLastId() + 1)
	
	newCoin := &Coin {
		Id: quant,
		Name: req.Name,
		Short: req.Short,
		Votes: 0,
	}
	
	_, err := MongoCollection.InsertOne(ctx, newCoin)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	return &pb.CryptoId{
		Id: quant,
	}, nil
}
