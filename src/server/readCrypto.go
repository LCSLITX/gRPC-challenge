package main

import(
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
	pb "github.com/lucassauro/klever-challenge/src/proto"
)

// ReadCrypto function receives context and request with id as parameters, return a crypto if it exists and an error.
func (s *Server) ReadCrypto(ctx context.Context, req *pb.CryptoId) (*pb.Crypto, error) {
	coin, err := DoesThisCryptoExist(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln(err))
	}
	
	return &pb.Crypto{
		Id: coin.ID,
		Name: coin.Name,
		Short: coin.Short,
		Votes: coin.Votes,
	}, nil
}