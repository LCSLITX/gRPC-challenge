package main

import (
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/lucassauro/klever-challenge/proto"
)

var address string = "0.0.0.0:50051"
type Server struct {
	pb.CryptoServiceServer
}

func main() {
	listener, err := net.Listen("tcp", address)
	
	if err != nil {
		status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	log.Println("Listening on port", address)

	serverInstance := grpc.NewServer()

	pb.RegisterCryptoServiceServer(serverInstance, &Server{})
	
	if err := serverInstance.Serve(listener); err != nil {
		status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}
}