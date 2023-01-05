package services

import (
	"baby-chain/gpp"
	"baby-chain/gpp/models"
	"baby-chain/tools"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func SendBC(ip string, buffer *bytes.Buffer) error {
	url := "http://" + ip + ":9090" + "/baby_chain/service/sync"
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

func Sync(ctx *gin.Context) {
	receiveObj := new(models.ReceiveSync)
	httpRes := gpp.NewHttpResponse(ctx)
	if err := httpRes.BindJson(&receiveObj); err != nil {
		httpRes.Error(err)
		return
	}

	sendObj := new(models.SendSync)
	data, err := json.Marshal(gpp.Bc)
	if err != nil {
		httpRes.Error(err)
		return
	}
	sendObj.BlockchainData = string(data)
	sendObj.Type = "receive"
	sendObjBytes, err := json.Marshal(sendObj)
	if err != nil {
		httpRes.Error(err)
		return
	}
	responseBody := bytes.NewBuffer(sendObjBytes)
	nodes, _ := gpp.Sd.Data["Nodes"].(tools.Data)
	for _, ip := range nodes {
		ip, _ := ip.(string)
		go func() {
			if err := SendBC(ip, responseBody); err != nil {
				httpRes.Error(err)
			}
		}()
	}

	httpRes.Text("ok")
}