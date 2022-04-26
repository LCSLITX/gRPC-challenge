package main

import (
	"io"
	"log"
	"context"
	
	"google.golang.org/protobuf/types/known/emptypb"
	pb "github.com/lucassauro/klever-challenge/src/proto"

)

func listCryptos(c pb.CryptoServiceClient) {
	stream, err := c.ListCryptos(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(res)
	}
}