package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/golang-api/server/routes"
	"log"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port: "5000",
		server: gin.Default(),
	}
}

func (server *Server) Run() {
	router := routes.ConfigurationRouter(server.server)

	log.Print("Server is running at port: ", server.port)
	log.Fatal(router.Run(":" + server.port))
}
