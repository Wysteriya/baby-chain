package services

import (
	"blockchain/blockchain"
	"github.com/gin-gonic/gin"
	"gpp/chain"
	"net/http"
)

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
		b := bch.BlockAt(bch.Len() - i)
		if err := cons.Exec(&bc, b); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if err := stts.Exec(&sd, b); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if err := bc.AddBlock(b); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}
	if err := chain.SaveBlockchain(&bc); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := chain.SaveStateData(&sd); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
}
