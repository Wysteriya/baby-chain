package gpp

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/consensus"
	"baby-chain/blockchain/state"
	"baby-chain/tools"
	"baby-chain/tools/data"
	"encoding/json"
	"errors"
	"os"
)

var Bc blockchain.Blockchain
var bcFName = "gpp/blockchain.json"
var Sd state.StateData
var sdFName = "gpp/statedata.json"
var Cons consensus.CAlgo
var States state.SAlgo

func genesisData() data.Data {
	return data.Data{}
}

func stateData() state.StateData {
	return state.StateData{Data: data.Data{}}
}

func FetchHyperParams() {
	if _, err := os.Open(bcFName); errors.Is(err, os.ErrNotExist) {
		Bc = blockchain.New(genesisData())
	} else {
		if err := json.Unmarshal(tools.ReadData(bcFName), &Bc); err != nil {
			panic(err)
		}
	}
	if _, err := os.Open(sdFName); errors.Is(err, os.ErrNotExist) {
		Sd = stateData()
	} else {
		if err := json.Unmarshal(tools.ReadData(sdFName), &Sd); err != nil {
			panic(err)
		}
	}

	Cons = consensus.New()
	States = state.New()
}

func SaveHyperParams() {
	save, err := json.Marshal(&Bc)
	if err != nil {
		panic(err)
	}
	if err := tools.WriteData(bcFName, save); err != nil {
		panic(err)
	}
	save, err = json.Marshal(&Sd)
	if err := tools.WriteData(sdFName, save); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}
