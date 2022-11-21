package main

import (
	"github.com/gin-gonic/gin"
	"gpp/client_services/services"
	"log"
)

func main() {
	server := gin.Default()
	basepath := server.Group("/rideshare")
	services.RegisterClientRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}
