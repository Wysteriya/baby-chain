package services

import (
	"baby-chain/blockchain/block"
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
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
	b, sendObj.PublicKey, sendObj.PrivateKey = gpp.Bc.MineNode(tools.Data{
		"driver":      receiveObj.Driver,
		"licenceNum":  receiveObj.LicenceNum,
		"vehicleType": receiveObj.VehicleType,

		"bloodGroup": receiveObj.BloodGroup,
		"age":        receiveObj.Age,
		"gender":     receiveObj.Gender,
	})
	if err := gpp.Cons.Exec(&gpp.Bc, b); err != nil {
		httpRes.Error(err)
		return
	}

	syncSendObj := new(models.SendSync)
	syncSendObj.Type = "send"
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(syncSendObj); err != nil {
		httpRes.Error(err)
		return
	}
	var err error
	ctx.Request, err = http.NewRequest(http.MethodPost, "localhost:9090/baby_chain/service/sync", &buf)
	if err != nil {
		httpRes.Error(err)
		return
	}
	httpRes.SendJson(sendObj)
	Sync(ctx)
}
