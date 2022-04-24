package main

import(
	"fmt"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/lucassauro/klever-challenge/src/proto"
)

// ListCryptos receives empty and stream as parameters. Returns a stream with all cryptos found in mongo collection.
func (s *Server) ListCryptos(no *empty.Empty, stream pb.CryptoService_ListCryptosServer) error {
	res, err := MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		status.Errorf(codes.Internal, fmt.Sprintln(err))
	}

	defer res.Close(context.Background())
	
	coin := &Coin{}

	for res.Next(context.Background()) {
		err := res.Decode(coin)

		if err != nil {
			status.Errorf(codes.Internal, fmt.Sprintln(err))
		}

		stream.Send(&pb.Crypto{
			Id: coin.ID,
			Name: coin.Name,
			Short: coin.Short,
			Votes: coin.Votes,
		})
	}

	if err := res.Err(); err != nil {
		status.Errorf(codes.Internal, fmt.Sprintln(err))
	}

	return nil
}
