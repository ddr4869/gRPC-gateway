package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/ddr4869/gRPC-gateway/proto/message"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) GetUnreadMessagesCount(ctx context.Context, req *pb.UnreadMessagesRequest) (*pb.UnreadMessagesResponse, error) {

	// make random int number
	unreadCount := rand.Int31()

	return &pb.UnreadMessagesResponse{
		UnreadCount: unreadCount,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server{})
	log.Printf("message-service server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
