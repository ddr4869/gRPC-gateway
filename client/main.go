package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ddr4869/gRPC-gateway/proto/message"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:50051"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect to message-service: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	userID := "user123"
	friendID := "friend456"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UnreadMessagesRequest{
		UserId:   userID,
		FriendId: friendID,
	}

	res, err := client.GetUnreadMessagesCount(ctx, req)
	if err != nil {
		log.Fatalf("could not get unread messages count: %v", err)
	}

	fmt.Printf("Unread messages count from %s to %s: %d\n", friendID, userID, res.UnreadCount)
}
