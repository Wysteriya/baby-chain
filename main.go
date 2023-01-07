package main

import (
	"baby-chain/gpp"
	"baby-chain/gpp/services"
	"github.com/gin-gonic/gin"
	"log"
)

func publicRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/public")
	clientRoute.GET("/sync", services.SyncGet)
	clientRoute.POST("/sync", services.SyncPost)
}

func privateRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/private")
	clientRoute.POST("/node", services.NodePost)
}

func main() {
	gpp.FetchHyperParams()
	chainName := "baby_chain"
	go func() {
		server := gin.Default()
		basePath := server.Group("/" + chainName)
		privateRoutes(basePath)
		log.Fatalln(server.Run("127.0.0.1:9080"))
	}()
	server := gin.Default()
	basePath := server.Group("/" + chainName)
	publicRoutes(basePath)
	log.Fatalln(server.Run(":9090"))
}
