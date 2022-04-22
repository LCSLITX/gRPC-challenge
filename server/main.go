package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "github.com/lucassauro/klever-challenge/proto"
)

var address string = "50051"

type Server struct {
	pb.CryptoServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error: %w", err)
	}

	fmt.Println("Listening on port: %w" , address)

	serverInstace := grpc.NewServer()
	if err := serverInstace.Serve(listener); err != nil {
		fmt.Println("Error: %w", err)
	}
	pb.RegisterCryptoServiceServer(serverInstace, &Server{})
}