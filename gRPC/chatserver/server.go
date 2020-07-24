package main

import (
	"github.com/kiranparajuli589/TourOfGo/gRPC/chatserver/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main()  {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listed on port 9000: %v", err)
	}

	s := chat.Server{}

	gRPCServer := grpc.NewServer()

	chat.RegisterChatServiceServer(gRPCServer, &s)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}


}
