package services

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/states"
	"baby-chain/gpp/chain"
	"fmt"
	"net/http"
)

import (
	"github.com/gin-gonic/gin"
)

var bc = chain.LoadBlockchain()
var sd = chain.LoadStateData()
var cons = chain.LoadConsensus()
var stts = chain.LoadStates()

func ExecSaveSync(ctx *gin.Context, b block.Block, bc *blockchain.Blockchain, sd *states.StateData) {
	httpResponse := newHttpResponse(ctx)
	if err := cons.Exec(bc, b); err != nil {
		httpResponse.Error(err)
		return
	}
	if err := stts.Exec(sd, b); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := chain.SaveBlockchain(bc); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := chain.SaveStateData(sd); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	go func() {
		if err := chain.Sync(bc, sd); err != nil {
			fmt.Println(err)
		}
	}()
}

func RegisterClientRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/service")

	clientRoute.POST("/newnode", NewNode)
	clientRoute.POST("/sync", Sync)
	clientRoute.POST("/publicinfo", PublicInfo)
	clientRoute.POST("/announcetravel", AnnounceTravel)
	clientRoute.POST("/messaging", MessageService)
	clientRoute.GET("/getopenannouncements", GetOpenAnnouncements)
	clientRoute.POST("/sendbid", SendBid)
}
