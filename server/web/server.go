package web

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine

	server   http.Server
	listener net.Listener
}

func (s *Server) Start() error {
	return nil
}
