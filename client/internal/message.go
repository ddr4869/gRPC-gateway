package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ddr4869/gRPC-gateway/proto/message"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetNumberOfUnreadMessage(c *gin.Context) {
	log.Print("1")
	address := "localhost:50051"

	// gRPC 연결 설정
	conn, err := grpc.NewClient(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Print("2")
	client := pb.NewMessageServiceClient(conn)
	log.Print("3")
	// 특정 유저와 친구 ID 설정
	userID := "user123"
	friendID := "friend456"

	// gRPC 요청 보내기
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Print("4")
	req := &pb.UnreadMessagesRequest{
		UserId:   userID,
		FriendId: friendID,
	}
	log.Print("5")
	res, err := client.GetUnreadMessagesCount(ctx, req)
	if err != nil {
		log.Fatalf("could not get unread messages count: %v", err)
	}
	// 결과 출력
	fmt.Printf("Unread messages count from %s to %s: %d\n", friendID, userID, res.UnreadCount)
}
