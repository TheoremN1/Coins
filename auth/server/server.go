package server

import (
	"fmt"

	"github.com/TheoremN1/Coins/auth/envs"
	"github.com/TheoremN1/Coins/auth/server/api"
	"github.com/gin-gonic/gin"
)

type Server struct {
	host   string
	port   string
	router *gin.Engine
}

var server *Server

func init() {
	router := api.GetRouter()
	env := envs.GetServerEnv()
	server = &Server{env.Host, env.Port, router}
}

func GetServer() *Server {
	return server
}

func (s *Server) Run() {
	s.router.Run(fmt.Sprintf("%s:%s", s.host, s.port))
}
