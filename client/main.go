package main

import(
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

var address string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error")
	}

	defer conn.Close()

	client := pb.NewCryptoServiceClient(conn)

	createCrypto(client)
	deleteCrypto(client, 1)
}