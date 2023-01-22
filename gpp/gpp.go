package gpp

import (
	"baby-chain/blockchain"
	cons "baby-chain/blockchain/consensus_state"
	"baby-chain/tools"
	"baby-chain/tools/data"
	"encoding/json"
	"errors"
	"os"
)

const (
	bcFName = "gpp/blockchain.json"
	sdFName = "gpp/statedata.json"
)

var (
	BC     blockchain.Blockchain
	SD     cons.StateData
	CSAlgo cons.CSAlgo
)

func genesisData() data.Data {
	return data.Data{}
}

func FetchHyperParams() (blockchain.Blockchain, cons.StateData, cons.CSAlgo) {
	var bc blockchain.Blockchain
	var sd cons.StateData
	var csAlgo = cons.New()
	if _, err := os.Open(bcFName); errors.Is(err, os.ErrNotExist) {
		bc = blockchain.New(genesisData())
	} else {
		if err := json.Unmarshal(tools.ReadData(bcFName), &bc); err != nil {
			panic(err)
		}
	}
	if _, err := os.Open(sdFName); errors.Is(err, os.ErrNotExist) {
		sd, err = csAlgo.GenSD(&bc)
		if err != nil {
			panic(err)
		}
	} else {
		if err := json.Unmarshal(tools.ReadData(sdFName), &sd); err != nil {
			panic(err)
		}
	}
	return bc, sd, csAlgo
}

func SaveHyperParams() {
	save, err := json.Marshal(&BC)
	if err != nil {
		panic(err)
	}
	if err := tools.WriteData(bcFName, save); err != nil {
		panic(err)
	}
	save, err = json.Marshal(&SD)
	if err := tools.WriteData(sdFName, save); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}
