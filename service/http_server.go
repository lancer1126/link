package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	_ server = (*httpServer)(nil)
)

type httpServer struct {
	*baseServer
	e      *gin.Engine
	server *http.Server
}

func (s *httpServer) start() error {
	return s.server.ListenAndServe()
}

func (s *httpServer) stop() error {
	return s.server.Shutdown(context.Background())
}
