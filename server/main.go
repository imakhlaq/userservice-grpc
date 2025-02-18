package main

import (
	"fmt"
	"net"

	"github.com/imakhlaq/userservice-grpc/handlers"
	pb "github.com/imakhlaq/userservice-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &handlers.Server{})
	fmt.Println("Server started at :", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to start: %v", err)
	}
}
