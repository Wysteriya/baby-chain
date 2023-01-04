package gpp

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/consensus"
	"baby-chain/blockchain/state"
	"baby-chain/tools"
	"encoding/json"
	"errors"
	"os"
)

var bc blockchain.Blockchain
var bcFName = "gpp/blockchain.json"
var sd state.StateData
var sdFName = "gpp/statedata.json"
var cons consensus.CAlgo
var states state.SAlgo

func genesisData() tools.Data {
	return tools.Data{}
}

func stateData() state.StateData {
	return state.StateData{Data: tools.Data{}}
}

func FetchHyperParams() {
	if _, err := os.Open(bcFName); errors.Is(err, os.ErrNotExist) {
		bc = blockchain.New(genesisData())
	} else {
		if err := json.Unmarshal(tools.ReadData(bcFName), &bc); err != nil {
			panic(err)
		}
	}
	if _, err := os.Open(sdFName); errors.Is(err, os.ErrNotExist) {
		sd = stateData()
	} else {
		if err := json.Unmarshal(tools.ReadData(sdFName), &sd); err != nil {
			panic(err)
		}
	}

	cons = consensus.New()
	states = state.New()
}

func SaveHyperParams() {
	save, err := json.Marshal(&bc)
	if err != nil {
		panic(err)
	}
	if err := tools.WriteData(bcFName, save); err != nil {
		panic(err)
	}
	save, err = json.Marshal(&sd)
	if err := tools.WriteData(sdFName, save); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}
