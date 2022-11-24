package main

import (
	"github.com/gin-gonic/gin"
	"gpp/services"
	"log"
)

func main() {
	chainName := "baby_chain"
	server := gin.Default()
	basePath := server.Group("/" + chainName)
	services.RegisterClientRoutes(basePath)
	log.Fatalln(server.Run(":9090"))
}
