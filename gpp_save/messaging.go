package gpp_save

import (
	"blockchain/block"
	"blockchain/wallet"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func MessageService(ctx *gin.Context) {
	responseObj := new(MessagePost)
	if err := ctx.BindJSON(&responseObj); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if responseObj.Type == "send" {
		toToPublicKey := responseObj.ToPublicKey
		nodes, _ := sd["full_nodes"].(block.Data)
		toIp, _ := nodes[toToPublicKey].(string)
		url := "http://" + toIp + ":9090/baby_chain/service/messaging"
		sign := wallet.SignMessage(responseObj.PrivateKeyOrSignature, []byte(responseObj.Message))
		sendObj := MessagePost{
			"receive",
			responseObj.PublicKey,
			string(sign),
			"",
			responseObj.Message,
		}
		send, err := json.Marshal(sendObj)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(send))
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		get, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.IndentedJSON(resp.StatusCode, string(get))
		return
	} else {
		msg := responseObj.Message
		wallet.VerifySignature(responseObj.PublicKey, []byte(msg), []byte(responseObj.PrivateKeyOrSignature))
		fmt.Println(msg)
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})
}
