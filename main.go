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
	//clientRoute.POST("/sync", Sync)
	//clientRoute.POST("/publicinfo", PublicInfo)
	//clientRoute.POST("/announcetravel", AnnounceTravel)
	//clientRoute.POST("/messaging", MessageService)
	//clientRoute.GET("/getopenannouncements", GetOpenAnnouncements)
	//clientRoute.POST("/sendbid", SendBid)
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
