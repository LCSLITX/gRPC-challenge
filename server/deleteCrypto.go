package main

import (
	"fmt"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/lucassauro/klever-challenge/proto"
)

func (s *Server) DeleteCrypto(ctx context.Context, req *pb.CryptoId) (*empty.Empty, error) {
	id := bson.M{"_id": req.Id}

	_, err := MongoCollection.DeleteOne(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("No coin with specified Id. %v", err))
	}

	return &empty.Empty{}, nil
}
