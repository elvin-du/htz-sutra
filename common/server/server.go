package server

import (
	"github.com/gin-gonic/gin"
	"strings"
)

var DefaultServer *Server = nil

func init() {
	DefaultServer = &Server{}
	DefaultServer.engine = gin.New()
}

type Server struct {
	engine *gin.Engine
}

func (s *Server) RegisterRoute(method, path string, handler gin.HandlerFunc) {
	switch strings.ToUpper(method) {
	case "GET":
		s.engine.GET(path, handler)
	case "POST":
		s.engine.POST(path, handler)
	default:
		panic("Unsupported http method:" + method)
	}
}

func (s *Server) Start(addr string ) {
	err := s.engine.Run(addr)
	if nil != err {
		panic(err)
	}
}
