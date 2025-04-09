package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Kpatoc452/cli_messanger/gen"
	"google.golang.org/grpc"
)

const port = 9000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	myServer := NewServer()
	myServer.Run(context.Background())
	grpcServer := grpc.NewServer()
	pb.RegisterChatServer(grpcServer, myServer)
	grpcServer.Serve(lis)
}
