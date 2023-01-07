package services

import (
	"baby-chain/blockchain"
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func SendIP(ip string, buffer *bytes.Buffer, service string) error {
	url := "http://" + ip + ":9090" + "/baby_chain/service/" + service
	resp, err := http.Post(url, "application/json", buffer)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		sb := string(body)
		return errors.New(sb)
	}
	return nil
}

func SendAll(httpRes *gpp.HttpResponse, sendObjBytes []byte, service string) {
	nodes, _ := gpp.Sd.Data["Nodes"].(tools.Data)
	for _, ip := range nodes {
		ip, _ := ip.(string)
		go func() {
			if err := SendIP(ip, bytes.NewBuffer(sendObjBytes), service); err != nil {
				httpRes.Error(err)
			}
		}()
	}
}

func Sync(ctx *gin.Context) {
	receiveObj := new(models.ReceiveSync)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	if receiveObj.Type == "send" {
		sendObj := new(models.SendSync)
		data, err := json.Marshal(gpp.Bc)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		sendObj.BlockchainData = string(data)
		sendObj.Type = "receive"
		sendObjBytes, err := json.Marshal(sendObj)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		SendAll(httpRes, sendObjBytes, "sync")
	} else {
		var bc blockchain.Blockchain
		if err := json.Unmarshal([]byte(receiveObj.BlockchainData), &bc); err != nil {
			httpRes.Error(err)
			return
		}
		if err := bc.Validate(); err != nil {
			httpRes.Error(err)
			return
		}
		if bc.Len() < gpp.Bc.Len() {
			httpRes.Error(fmt.Errorf("outdatedBlockchainReceived"))
			return
		}
		if bc.Chain[gpp.Bc.Len()-1].Hash != gpp.Bc.CurrHash() {
			httpRes.Error(fmt.Errorf("blockchainCompatibilityError"))
			return
		}
		gpp.Bc = bc
	}

	log.Println("sync ok")
}
