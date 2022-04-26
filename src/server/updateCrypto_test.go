package main

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

func Test_UpdateCrypto(t *testing.T) {
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

	
	mt.Run("Success", func(mt *mtest.T) {
		res, _ := client.CreateCrypto(ctx, &pb.NewCrypto{
			Name: "TestUpdateCrypto",
			Short: "TVT",
		})

		mt.AddMockResponses(bson.D{
			primitive.E	{ Key: "ok", Value: 1 },
			primitive.E {
				Key: "value", 
				Value: bson.D {
					primitive.E{ Key: "Name", Value: "Teste"},
					primitive.E{ Key: "Short", Value: "TST"},
			}}})

		_, err := client.UpdateCrypto(context.Background(), &pb.Crypto{
			Id: res.Id,
			Name: "Teste",
			Short: "TST",
		})

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}
