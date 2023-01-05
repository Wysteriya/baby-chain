package main

import (
	"baby-chain/gpp"
	"baby-chain/gpp/services"
	"github.com/gin-gonic/gin"
	"log"
)

func registerClientRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/service")

	clientRoute.POST("/node", services.NodePost)
	clientRoute.POST("/announcetravel", services.AnnounceTravel)
	clientRoute.POST("/getannouncetravel", services.GetAnnounceTravel)
	clientRoute.POST("/sendbid", services.SendBid)
	clientRoute.POST("/getbid", services.GetBid)
	//clientRoute.POST("/messaging", )
	//clientRoute.GET("/getopenannouncements", )
}

func main() {
	gpp.FetchHyperParams()
	chainName := "baby_chain"
	server := gin.Default()
	basePath := server.Group("/" + chainName)
	registerClientRoutes(basePath)
	log.Fatalln(server.Run(":9090"))
	//gpp.SaveHyperParams()
}
