package services

import (
	"baby-chain/blockchain/block"
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools"
	"github.com/gin-gonic/gin"
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

	httpRes.SendJson(sendObj)
}
