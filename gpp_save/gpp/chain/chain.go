package chain

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/consensus"
	"baby-chain/blockchain/states"
	"baby-chain/tools"
	"encoding/json"
)

import (
	"fmt"
	"os"
)

var bcFile = "../blockchain/db/blockchain.bin"
var sdFile = "../blockchain/db/statedata.bin"

func LoadBlockchain() blockchain.Blockchain {
	bchData := tools.ReadData(bcFile)
	//bc, err := blockchain.Load(bchData)
	var bc blockchain.Blockchain
	err := json.Unmarshal(bchData, &bc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bc
}

func SaveBlockchain(bc *blockchain.Blockchain) error {
	save, err := bc.MarshalJSON()
	if err != nil {
		return err
	}
	if err := tools.WriteData(bcFile, save); err != nil {
		return err
	}
	return nil
}

func LoadStateData() states.StateData {
	sdData := tools.ReadData(sdFile)
	var sd states.StateData
	err := json.Unmarshal(sdData, &sd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return sd
}

func SaveStateData(sd *states.StateData) error {
	save, err := json.Marshal(sd)
	if err != nil {
		return err
	}
	if err := tools.WriteData(sdFile, save); err != nil {
		return err
	}
	return nil
}

func LoadConsensus() consensus.CAlgo {
	return consensus.New(CPublicInfo, CAnnounceTravel)
}

func LoadStates() states.SAlgo {
	return states.New(SPublicInfo, SAts)
}
