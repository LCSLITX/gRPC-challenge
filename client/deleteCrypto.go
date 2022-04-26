package main

import(
	"log"
	"context"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

func deleteCrypto(client pb.CryptoServiceClient, id uint32) uint32 {
	if _, err := client.DeleteCrypto(context.Background(), &pb.CryptoId{
		Id: id,
	}); err != nil {
		log.Fatalln(err)
	}

	log.Println(id)

	return id
}