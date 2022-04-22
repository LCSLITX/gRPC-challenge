package main

import (
	"log"
	"net"

	pb "github.com/lucassauro/klever-challenge/proto"
	"google.golang.org/grpc"
)

var address string = ":50051"

type Server struct {
	pb.CryptoServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return 
	}
	log.Println("Listening on port", address)

	serverInstace := grpc.NewServer()
	if err := serverInstace.Serve(listener); err != nil {
		log.Fatalf("Error: %v", err)
	}
	pb.RegisterCryptoServiceServer(serverInstace, &Server{})
}