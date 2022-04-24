package main

import (
	"context"

	pb "github.com/lucassauro/klever-challenge/proto"
)

// ReadCrypto function receives context and request with id as parameters, return a crypto if it exists and an error.
func (s *Server) ReadCrypto(ctx context.Context, req *pb.CryptoId) (*pb.Crypto, error) {
	
	coin := DoesThisCryptoExist(req.Id)

	return &pb.Crypto{
		Id: coin.Id,
		Name: coin.Name,
		Short: coin.Short,
		Votes: coin.Votes,
	}, nil
}