package grpc

import (
	"context"

	pb "github.com/ddr4869/gRPC-gateway/proto/message"
	"google.golang.org/grpc"
)

var (
	GrpcClient grpc.ClientConn
)

func NewClient(address string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.NewClient(address, opts...)
}

type Server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *Server) GetUnreadMessagesCount(ctx context.Context, req *pb.UnreadMessagesRequest) (*pb.UnreadMessagesResponse, error) {
	_ = req.GetUserId()
	_ = req.GetFriendId()

	// 여기에 DB 조회 로직을 추가합니다.
	// 예시로 하드코딩된 값을 반환합니다.
	unreadCount := int32(5)

	return &pb.UnreadMessagesResponse{
		UnreadCount: unreadCount,
	}, nil
}
