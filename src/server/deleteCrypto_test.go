package main

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/credentials/insecure"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

func Test_Error_DeleteCrypto(t *testing.T) {
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
		mt.AddMockResponses(bson.D{
			{Key: "DeletedCount", Value: 0},
		})
		_, err := client.DeleteCrypto(ctx, &pb.CryptoId{
			Id: 1,
		})
		
		if err == nil {
			t.Errorf("Error expected")
		}

		stat, ok := status.FromError(err)

		if !ok {
			t.Errorf("Error expected")
		}

		if stat.Code() != codes.NotFound {
			t.Errorf("Expected NotFound, got %v", stat.Code().String())
		}
	})
}

func Test_Error_DeleteInvalidID(t *testing.T) {
	ctx := context.Background()
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), credentials)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer connection.Close()
	client := pb.NewCryptoServiceClient(connection)

	_, err = client.DeleteCrypto(context.Background(), &pb.CryptoId{})

	if err == nil {
		t.Error("Error expected")
	}

	stat, ok := status.FromError(err)

	if !ok {
		t.Error("Error expected")
	}

	if stat.Code() != codes.NotFound {
		t.Errorf("Expected NotFound, got %v", stat.Code().String())
	}
}