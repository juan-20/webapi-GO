package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/juan-20/webapi-GO/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "300",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("Server is running in port ")
	log.Fatal(router.Run((":" + s.port)))
}
