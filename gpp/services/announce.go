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

func GetAnnounceTravel(ctx *gin.Context) {
	receiveObj := new(models.SendAnnouncement)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	var b block.Block
	if err := json.Unmarshal([]byte(receiveObj.AnnounceBlockData), &b); err != nil {
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

func AnnounceTravel(ctx *gin.Context) {
	receiveObj := new(models.ReceiveAnnouncement)
	sendObj := new(models.SendAnnouncement)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	b := gpp.Bc.MineBlock("AnnounceTravel",
		tools.Data{
			"public_key": receiveObj.PublicKey,
			"from_lat":   receiveObj.FromLat,
			"from_lon":   receiveObj.FromLon,
			"to_lat":     receiveObj.ToLat,
			"to_lon":     receiveObj.ToLon,
			"time":       receiveObj.Time,
		},
	)
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
	announceBlockData, err := json.Marshal(&b)
	if err != nil {
		httpRes.Error(err)
		return
	}
	sendObj.AnnounceBlockData = string(announceBlockData)
	sendObjBytes, err := json.Marshal(sendObj)
	if err != nil {
		httpRes.Error(err)
		return
	}
	SendAll(httpRes, bytes.NewBuffer(sendObjBytes), "getannouncetravel") // todo: sync to drivers only

	httpRes.Text("ok")
}
