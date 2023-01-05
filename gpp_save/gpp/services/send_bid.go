package services

import (
	"baby-chain/blockchain/wallet"
	"baby-chain/tools"
	"github.com/gin-gonic/gin"
)

func SendBid(ctx *gin.Context) {
	responseObj := new(SendBidPost)

	httpResponse := newHttpResponse(ctx)
	httpResponse.BindJson(responseObj)

	announcementId := responseObj.AnnouncementId
	publicKey := responseObj.PublicKey
	privateKey := responseObj.PrivateKey
	amount := responseObj.BidAmount

	b := bc.MineBlock(
		"Bid",
		tools.Data{"public_key": publicKey, "bid_amount": amount, "announcement_id": announcementId},
	)
	hash := b.Hash
	signature, err := wallet.SignHash(privateKey, tools.HashB(hash[:]))
	if err != nil {
		httpResponse.Error(err)
	}
	b.Header["signature"] = string(signature)
	ExecSaveSync(ctx, b, &bc, &sd)
	httpResponse.Text("sent bid successfully")
}
