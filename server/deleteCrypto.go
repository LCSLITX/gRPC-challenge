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

// DeleteCrypto function receives context and request with crypto id. Returns empty in case of success, error in case of failure.
func (s *Server) DeleteCrypto(ctx context.Context, req *pb.CryptoId) (*empty.Empty, error) {

	id := bson.M{"_id": req.Id}

	res, err := MongoCollection.DeleteOne(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln(err))
	} else if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln("DeletedCount", res.DeletedCount))
	}

	return &empty.Empty{}, nil
}
