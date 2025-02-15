package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetUp(s *Server) {
	r := s.router
	api := r.Group("/api/v1")
	api.GET("/ping", s.Ping)
	api.GET("/message", s.GetNumberOfUnreadMessage)

}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func (s *Server) Start() error {

	srv := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: s.router,
	}

	log.Printf("Listening and serving HTTP on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("listen: %s\n", err)
		return err
	}

	return nil
}
