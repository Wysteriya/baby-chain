package chain

import (
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/states"
	"baby-chain/tools"
)

func initializeSBS(sd *states.StateData) {
	(*sd).Data["bids"] = tools.Data{}
}

func validateSBS(_ *states.StateData, b block.Block) bool {
	header := b.Header
	if header["head"] != "Bid" {
		return false
	}
	return true
}

func runSBS(sd *states.StateData, b block.Block) error {
	data := b.Data
	if (*sd).Data["bids"] == nil {
		initializeSBS(sd)
	}
	oa, _ := (*sd).Data["bids"].(tools.Data)
	key := tools.HashB([]byte(data.String())).Hex()
	oa[key] = b.Data
	return nil
}

var SSBs = states.State{Check: validateSBS, Run: runSBS}
