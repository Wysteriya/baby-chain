package main

import (
	"github.com/gin-gonic/gin"
	"gpp/services"
	"log"
)

func main() {

	server := gin.Default()
	basepath := server.Group("/rideshare")
	services.RegisterClientRoutes(basepath)
	services.RegisterBlockchainRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}
