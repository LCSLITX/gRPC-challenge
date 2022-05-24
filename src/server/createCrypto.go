package main

import(
	"fmt"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-playground/validator/v10"
	pb "github.com/lucassauro/klever-challenge/src/proto"
)

// CreateCrypto receives context and request with NewCrypto struct as parameters, 
// create a coin in the database and return its auto incremented integer Id.
func (s *Server) CreateCrypto(ctx context.Context, req *pb.NewCrypto) (*pb.CryptoId, error) {
	lastID, Err := FindLastID()

	if Err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln(Err))
	}
  // commit teste
	// Following variable stores the Id of the last inserted document plus one.
	generatedCoinID := lastID + 1
	
	newCoin := &Coin {
		ID: generatedCoinID,
		Name: req.Name,
		Short: req.Short,
		Votes: 0,
	}

	// validate existence of both name and short 
	Validate = validator.New()
	
	if validationErr := Validate.Struct(newCoin);  validationErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintln(validationErr))
	}
	
	if _, err := MongoCollection.InsertOne(ctx, newCoin); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln(err))
	}

	return &pb.CryptoId{
		Id: generatedCoinID,
	}, nil
}
