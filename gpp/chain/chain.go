package chain

import (
	"blockchain/blockchain"
	"blockchain/consensus"
	"blockchain/db/jsoner"
	"blockchain/states"
)

import (
	"fmt"
	"os"
)

func LoadBlockchain() blockchain.Blockchain {
	bchData, err := jsoner.ReadData("../blockchain/db/blockchain.bin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bc, err := blockchain.Load(bchData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bc
}

func LoadStateData() states.StateData {
	sdData, err := jsoner.ReadData("../blockchain/db/statedata.bin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sd, err := states.Load(sdData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return sd
}

func LoadConsensus() consensus.CAlgo {
	return consensus.New()
}

func LoadStates() states.States {
	return states.New()
}
