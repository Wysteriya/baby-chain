package services

import (
	"baby-chain/blockchain/wallet"
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AnnounceTravel(ctx *gin.Context) {
	receiveObj := new(models.ReceiveAnnouncement)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}
	b := gpp.Bc.MineBlock(
		"AnnounceTravel",
		tools.Data{
			"public_key": receiveObj.PublicKey,
			"from_lat":   receiveObj.FromLat,
			"from_lon":   receiveObj.FromLon,
			"to_lat":     receiveObj.ToLat,
			"to_lon":     receiveObj.ToLon,
			"time":       receiveObj.Time,
		},
	)
	hash := b.Hash
	signature, err := wallet.SignHash(receiveObj.PrivateKey, tools.HashB(hash[:]))
	if err != nil {

	}
	b.Header["signature"] = string(signature)
	if err := gpp.Cons.Exec(&gpp.Bc, b); err != nil {
		fmt.Println("hi", err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := gpp.States.Exec(&gpp.Sd, b); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	gpp.SaveHyperParams()

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})

}
