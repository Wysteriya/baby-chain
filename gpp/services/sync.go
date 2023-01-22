package services

import (
	"baby-chain/blockchain"
	"baby-chain/gpp"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func SendIP(ip string, buffer *bytes.Buffer, service string) error {
	url := "http://" + ip + ":9090" + "/baby_chain/public/" + service
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

func SendAll(data []byte, service string) chan error {
	nodes, _ := gpp.Sd.Data["Nodes"].(data.Data)
	errChan := make(chan error)
	go func() {
		for _, ip := range nodes {
			ip, _ := ip.(string)
			if err := SendIP(ip, bytes.NewBuffer(data), service); err != nil {
				errChan <- err
			}
		}
		close(errChan)
	}()
	return errChan
}

func SyncSend() error {
	sendObjBytes, err := json.Marshal(gpp.Bc)
	if err != nil {
		return err
	}
	SendAll(sendObjBytes, "sync")
	return nil
}

func SyncPost(ctx *gin.Context) {
	var bc blockchain.Blockchain
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&bc); err != nil {
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

	httpRes.Text("ok")
}

func SyncGet(ctx *gin.Context) {
	httpRes := gpp.NewHttpResponse(ctx)
	httpRes.SendJson(gpp.Bc)
}
