package main

import(
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	pb "github.com/lucassauro/klever-challenge/src/proto"
)

const bufSize = 1024 * 1024

var listener *bufconn.Listener

func init() {
	listener = bufconn.Listen(bufSize)

	serverInstance := grpc.NewServer()
	
	pb.RegisterCryptoServiceServer(serverInstance, &Server{})

	go func() {
		if err := serverInstance.Serve(listener); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return listener.Dial()
}
