package chain

import (
	"blockchain/block"
	"blockchain/blockchain"
	"blockchain/states"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func Sync(bc *blockchain.Blockchain, sd *states.StateData) error {
	nodes, _ := (*sd)["full_nodes"].(block.Data)
	for _, ip := range nodes {
		ip, _ := ip.(string)
		save, err := bc.Save()
		if err != nil {
			return err
		}
		if err := SendBC(ip, save); err != nil {
			return err
		}
	}
	return nil
}

func SendBC(ip string, data []byte) error {
	dataString := string(data)
	postBody, _ := json.Marshal(map[string]string{
		"blockchain_data": dataString,
	})
	responseBody := bytes.NewBuffer(postBody)
	url := "http://" + ip + ":9090" + "/baby_chain/service/sync"
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		return errors.New(sb)
	}
	return nil
}
