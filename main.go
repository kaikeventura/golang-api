package main

import (
	"github.com/kaikeventura/golang-api/database"
	"github.com/kaikeventura/golang-api/server"
)

func main() {
	database.StartDatabase()
	ginServer := server.NewServer()

	ginServer.Run()
}
