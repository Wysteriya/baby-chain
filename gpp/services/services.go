package services

import (
	"blockchain/blockchain"
	"gpp/chain"
	"net/http"
)

import (
	"blockchain/block"
	"blockchain/wallet"
)

import (
	"github.com/gin-gonic/gin"
)

var bc = chain.LoadBlockchain()
var sd = chain.LoadStateData()
var cons = chain.LoadConsensus()
var stts = chain.LoadStates()

func NewNode(ctx *gin.Context) {
	responseObj := new(NewNodePost)
	if err := ctx.BindJSON(responseObj); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ipAddress := responseObj.IpAddress

	returnObj := new(NewNodeResponse)
	publicKey, privateKey, err := wallet.GeneratePublicAddressAndKey()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	returnObj.PublicKey = publicKey
	returnObj.PrivateKey = privateKey

	b := bc.MineBlock(
		block.Data{"head": "NewNode"},
		block.Data{"public_key": publicKey, "ip_address": ipAddress},
	)
	hash := b.Hash()
	signature, err := wallet.SignMessage(privateKey, publicKey, hash[:])
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	b.Header["signature"] = string(signature)
	if err := cons.Exec(&bc, b); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := stts.Exec(&sd, b); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := chain.Sync(&bc, &sd); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	bc.Print()
	if err := chain.SaveBlockchain(&bc); err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, returnObj)
}

func Sync(ctx *gin.Context) {
	responseObj := new(SyncPost)
	if err := ctx.BindJSON(responseObj); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	blockchainData := responseObj.BlockchainData
	bch, err := blockchain.Load([]byte(blockchainData))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if bc.Len() > bch.Len() {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "short length of blockchain"})
		return
	}
	diffLen := bch.Len() - bc.Len()
	if bch.HashOf(bch.Len()-diffLen-1) != bc.HashOf(bc.Len()-1) {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "hash mismatch"})
		return
	}
	for i := diffLen; i > 0; i-- {
		if err := bc.AddBlock(bch.BlockAt(bch.Len() - i)); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}
	if err := chain.SaveBlockchain(&bc); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
}

func RegisterClientRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/service")
	clientRoute.POST("/newnode", NewNode)
	clientRoute.POST("/Sync", Sync)
}
