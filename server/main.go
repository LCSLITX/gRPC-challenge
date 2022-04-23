package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/lucassauro/klever-challenge/proto"
)

var address string = "0.0.0.0:50051"
type Server struct {
	pb.CryptoServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)
	
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Println("Listening on port", address)

	serverInstance := grpc.NewServer()

	pb.RegisterCryptoServiceServer(serverInstance, &Server{})
	
	if err := serverInstance.Serve(listener); err != nil {
		log.Fatalf("Error: %v", err)
	}
}