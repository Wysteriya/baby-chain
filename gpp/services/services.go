package services

import (
	"fmt"
	"net/http"
	"os"
)

import (
	"blockchain/block"
	"blockchain/blockchain"
	"blockchain/consensus"
	"blockchain/db/jsoner"
	"blockchain/wallet"
)

import (
	"github.com/gin-gonic/gin"
)

func LoadBlockchain() blockchain.Blockchain {
	bchData, err := jsoner.ReadData("../blockchain.bin")
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
	ctx.IndentedJSON(http.StatusOK, returnObj)
}

func RegisterClientRoutes(rg *gin.RouterGroup) {
	clientRoute := rg.Group("/service")
	clientRoute.POST("/newnode", NewNode)
}
