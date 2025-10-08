package main

import (
	"fmt"
	"log"
	"net"
	server2 "test_task/internal/server"

	"google.golang.org/grpc"
	pb "test_task/pb/storage"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	service := server2.NewStorageService()
	pb.RegisterStorageServiceServer(server, service)

	fmt.Println("✅ gRPC server started on port 50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
