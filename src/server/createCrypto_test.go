package main

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

func TestCreateCrypto(t *testing.T) {
	ctx := context.Background()
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), credentials)
	// https://stackoverflow.com/questions/42102496/testing-a-grpc-service
	
	if err != nil {
	t.Fatalf("Error: %v", err)
	}

	defer connection.Close()

	client := pb.NewCryptoServiceClient(connection)
	
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	defer mt.Close()

	mt.Run("Succeeded", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		if _, err := client.CreateCrypto(
			ctx, 
			&pb.NewCrypto{
				Name: "Tests",
				Short: "Tst",
			}); err != nil {
				t.Errorf("Error: %v", err)
			}
	})
}

func TestErrorCreateCrypto(t *testing.T) {
	ctx := context.Background()
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), credentials)
	
	if err != nil {
	t.Fatalf("Error: %v", err)
	}

	defer connection.Close()

	client := pb.NewCryptoServiceClient(connection)
	
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	defer mt.Close()

	mt.Run("Failure", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D { primitive.E { Key: "error", Value: 0 } })

		_, err := client.CreateCrypto(ctx, &pb.NewCrypto{})

		if err == nil {
			t.Errorf("Error expected")
		}

		stat, ok := status.FromError(err)

		if !ok {
			t.Error("Error expected")
		}

		if stat.Code() != codes.InvalidArgument {
			t.Errorf("Expected InvalidArgument, got %v", stat.Code().String())
		}
	})
}