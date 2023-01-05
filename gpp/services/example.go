package services

import (
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"github.com/gin-gonic/gin"
)

func Example(ctx *gin.Context) {
	receiveObj := new(models.ReceiveModel)
	sendObj := new(models.SendModel)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}
	// custom code
	httpRes.SendJson(sendObj)
}
