package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	Router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{
		Router: router,
	}
}
func (s *Server) Run(address string) {
	log.Fatal(http.ListenAndServe(address, s.Router))
}
