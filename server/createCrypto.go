package main

import (
	"fmt"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-playground/validator/v10"
	pb "github.com/lucassauro/klever-challenge/proto"
)

// CreateCrypto receives context and request with NewCrypto struct as parameters, create a coin in the database and return its auto incremented integer Id.
func (s *Server) CreateCrypto(ctx context.Context, req *pb.NewCrypto) (*pb.CryptoId, error) {
	
	lastId, Err := FindLastId()

	if Err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln(Err))
	}

	// Following variable stores the Id of the last inserted document plus one.
	generatedCoinId := uint32(lastId + 1)
	
	newCoin := &Coin {
		Id: generatedCoinId,
		Name: req.Name,
		Short: req.Short,
		Votes: 0,
	}

	// validate existance of both name and short 
	validate = validator.New()
	
	validationErr := validate.Struct(newCoin)
	
	if validationErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln(validationErr))
	}
	
	_, err := MongoCollection.InsertOne(ctx, newCoin)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln(err))
	}

	return &pb.CryptoId{
		Id: generatedCoinId,
	}, nil
}
