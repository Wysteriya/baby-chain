package services

import (
	"baby-chain/blockchain"
	cons "baby-chain/blockchain/consensus_state"
	errors2 "baby-chain/errors"
	"baby-chain/gpp"
	. "baby-chain/tools/data"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func SendToIP(ip string, buffer *bytes.Buffer, service string) error {
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
	nodes, _ := gpp.SD.Data[cons.NODES].(Data)
	errChan := make(chan error)
	go func() {
		for _, ip := range nodes {
			ip, _ := ip.(string)
			if err := SendToIP(ip, bytes.NewBuffer(data), service); err != nil {
				errChan <- err
			}
		}
		close(errChan)
	}()
	return errChan
}

func SyncSend() error {
	sendObjBytes, err := json.Marshal(gpp.BC)
	if err != nil {
		return err
	}
	SendAll(sendObjBytes, "sync")
	return nil
}

func SyncPost(ctx *gin.Context) {
	var bc blockchain.Blockchain
	var sd cons.StateData
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&bc); err != nil {
		httpRes.Error(err)
		return
	}

	if bc.Len() < gpp.BC.Len() {
		httpRes.Error(errors2.OutdatedBlockchainReceived)
		return
	}
	if err := bc.Validate(); err != nil {
		httpRes.Error(err)
		return
	}
	if bc.Chain[gpp.BC.Len()-1].Hash != gpp.BC.CurrHash() {
		httpRes.Error(errors2.BlockchainCompatibilityError)
		return
	}
	sd, err := gpp.CSAlgo.SetSD(&bc, gpp.SD, gpp.BC.Len()-1, bc.Len())
	if err != nil {
		httpRes.Error(err)
		return
	}
	gpp.BC = bc
	gpp.SD = sd

	httpRes.Text("ok")
}

func SyncGet(ctx *gin.Context) {
	httpRes := gpp.NewHttpResponse(ctx)
	httpRes.SendJson(gpp.BC)
}
