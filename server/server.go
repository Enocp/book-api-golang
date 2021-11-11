package server

import (
	"github/Enocp/webapi-go/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server {
		port: "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Println("Server is runing at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}