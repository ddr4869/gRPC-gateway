package internal

import (
	"context"
	"net/http"

	pb "github.com/ddr4869/gRPC-gateway/proto/helloworld"
	"github.com/gin-gonic/gin"
)

func (s *Server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func (s *Server) GetNumberOfUnreadMessage(c *gin.Context) {

	c.JSON(http.StatusOK, "1")
}
