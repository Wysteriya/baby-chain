package services

import (
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/consensus"
	"baby-chain/blockchain/wallet"
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GetBid(ctx *gin.Context) {
	receiveObj := new(models.SendBid)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	var b block.Block
	if err := json.Unmarshal([]byte(receiveObj.BidBlockData), &b); err != nil {
		httpRes.Error(err)
		return
	}
	if err := consensus.SignCheckBlock(b); err != nil {
		httpRes.Error(err)
		return
	}
	if err := gpp.Cons.Exec(&gpp.Bc, b); err != nil {
		httpRes.Error(err)
		return
	}
	if err := gpp.States.Exec(&gpp.Sd, b); err != nil {
		httpRes.Error(err)
		return
	}

	httpRes.Text("ok")
}

func SendBid(ctx *gin.Context) {
	receiveObj := new(models.ReceiveBid)
	sendObj := new(models.SendBid)
	httpRes := gpp.NewHttpResponse(ctx)
	err := httpRes.BindJson(receiveObj)
	if err != nil {
		httpRes.Error(err)
		return
	}

	openAnnouncements := gpp.Sd.Data[receiveObj.AnnouncementId].(tools.Data)
	publicKey := openAnnouncements["public_key"].(string)
	nodes := gpp.Sd.Data["nodes"].(tools.Data)
	ipAddress := nodes[publicKey].(string)
	b := gpp.Bc.MineBlock("bids",
		tools.Data{"announcement_id": receiveObj.AnnouncementId,
			"bid_amount": receiveObj.BidAmount,
			"public_key": receiveObj.PublicKey})
	if err := gpp.Cons.Exec(&gpp.Bc, b); err != nil {
		httpRes.Error(err)
		return
	}
	if err := gpp.States.Exec(&gpp.Sd, b); err != nil {
		httpRes.Error(err)
		return
	}
	signature, err := wallet.SignHash(receiveObj.PrivateKey, b.Hash)
	if err != nil {
		httpRes.Error(err)
		return
	}
	b.Header["signature1"] = string(signature)
	bidBlocData, err := json.Marshal(&b)
	if err != nil {
		httpRes.Error(err)
		return
	}
	sendObj.BidBlockData = string(bidBlocData)
	sendObjBytes, err := json.Marshal(sendObj)
	if err != nil {
		httpRes.Error(err)
		return
	}
	err = SendIP(ipAddress, bytes.NewBuffer(sendObjBytes), "getbid")
	if err != nil {
		httpRes.Error(err)
		return
	}

	httpRes.Text("ok")

}
