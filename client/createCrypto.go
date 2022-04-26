package main

import(
	"log"
	"context"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

func createCrypto(client pb.CryptoServiceClient) uint32 {
	data := &pb.NewCrypto{
		Name: "Bitcoin",
		Short: "BTC",
	}

	res, err := client.CreateCrypto(context.Background(), data)

	if err != nil {
		log.Fatalf("Error")
	}

	return res.Id
}