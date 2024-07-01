package internal

import (
	"log"

	pb "github.com/ddr4869/gRPC-gateway/server/proto"
	"github.com/gin-gonic/gin"
)

func (s *Server) CallProductRpc(c *gin.Context) {
	product := &pb.Product{
		Id:    "A",
		Name:  "B",
		Price: 100,
	}
	log.Print(product)
}
