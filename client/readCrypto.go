package main

import(
	"log"
	"context"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

func readCrypto(client pb.CryptoServiceClient, id uint32) *pb.Crypto {
	data := &pb.CryptoId{
		Id: id,
	}

	res, err := client.ReadCrypto(context.Background(), data)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res)

	return res
}