package main

import(
	"log"
	"context"

	pb "github.com/lucassauro/klever-challenge/src/proto"
)

func updateCrypto(client pb.CryptoServiceClient, id uint32) uint32 {
	data := &pb.Crypto{
		Id: id,
		Name: "Bitcoin",
		Short: "BTC",
	}

	res, err := client.UpdateCrypto(context.Background(), data)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Id)

	return res.Id
}