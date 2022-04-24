package main

import (
	"context"
	"fmt"

	pb "github.com/lucassauro/klever-challenge/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReadCrypto function receives context and request with id as parameters, return a crypto if it exists and an error.
func (s *Server) ReadCrypto(ctx context.Context, req *pb.CryptoId) (*pb.Crypto, error) {
	coin, err := DoesThisCryptoExist(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln(err))
	}
	
	return &pb.Crypto{
		Id: coin.Id,
		Name: coin.Name,
		Short: coin.Short,
		Votes: coin.Votes,
	}, nil
}