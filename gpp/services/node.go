package services

import (
	"baby-chain/blockchain/block"
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"github.com/gin-gonic/gin"
	"log"
)

func NodePost(ctx *gin.Context) {
	receiveObj := new(models.ReceiveNode)
	sendObj := new(models.SendNode)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	var b block.Block
	b, sendObj.PublicKey, sendObj.PrivateKey = gpp.BC.MineNode(receiveObj.Data)
	if err := gpp.CSAlgo.Exec(&gpp.BC, &gpp.SD, b); err != nil {
		httpRes.Error(err)
		return
	}

	if err := SyncSend(); err != nil {
		log.Printf("sync failed: %s", err)
	}
	httpRes.SendJson(sendObj)
}
