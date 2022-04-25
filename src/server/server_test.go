package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lucassauro/klever-challenge/src/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)

	s := grpc.NewServer()
	pb.RegisterCryptoServiceServer(s, &Server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
