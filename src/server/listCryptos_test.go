package main

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/credentials/insecure"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

func Test_listCrypto(t *testing.T) {
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
		client.CreateCrypto(ctx, &pb.NewCrypto{
			Name: "TesteListCryptos",
			Short: "TST",
		})

		first := mtest.CreateCursorResponse(
			1,
			"cryptodb.crypto",
			mtest.FirstBatch,
			bson.D{ 
				primitive.E { Key: "name", Value: "Teste"},
				primitive.E { Key: "short", Value: "TST"},
			},
		)
		killCursors := mtest.CreateCursorResponse(
			0, 
			"cryptodb.crypto", 
			mtest.NextBatch,
		)
		mt.AddMockResponses(first, killCursors)
		// 
		_, err := client.ListCryptos(ctx, &empty.Empty{})

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}
