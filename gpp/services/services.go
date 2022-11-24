package services

import (
	"blockchain/block"
	"blockchain/blockchain"
	"blockchain/consensus"
	"blockchain/wallet"
	"db/jsoner"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func LoadBlockchain() blockchain.Blockchain {
	bchData, err := jsoner.ReadData("../db/blockchain.bin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bc, err := blockchain.Load(bchData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bc
}

var bc = LoadBlockchain()
var cons = consensus.New()

func NewNode(ctx *gin.Context) {
	responseObj := new(NewNodePost)
	if err := ctx.BindJSON(responseObj); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ipAddress := responseObj.IpAddress

	returnObj := new(NewNodeResponse)
	publicAddress, privateKey := wallet.GeneratePublicAddressAndKey()
	returnObj.PublicAddress = publicAddress
	returnObj.PrivateKey = privateKey

	b := bc.MineBlock(
		block.Data{"head": "NewNode"},
		block.Data{"public_address": publicAddress, "ip_address": ipAddress},
	)
	hash := b.Hash()
	signature, err := wallet.SignMessage(privateKey, publicAddress, hash[:])
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	b.Header["signature"] = string(signature)
	if err := cons.Exec(&bc, b); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, returnObj)
}

func RegisterClientRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/service")
	clientRoute.POST("/newnode", NewNode)
}
