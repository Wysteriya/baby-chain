package chain

import (
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/states"
	"baby-chain/tools"
)

func initializeATS(sd *states.StateData) {
	(*sd).Data["open_announcements"] = tools.Data{}
}

func validateATS(_ *states.StateData, b block.Block) bool {
	header := b.Header
	if header["head"] != "AnnounceTravel" {
		return false
	}
	return true
}

func runATS(sd *states.StateData, b block.Block) error {
	data := b.Data
	if (*sd).Data["open_announcements"] == nil {
		initializeATS(sd)
	}
	oa, _ := (*sd).Data["open_announcements"].(tools.Data)
	key := tools.HashB([]byte(data.String())).Hex()
	oa[key] = data
	return nil
}

var SAts = states.State{Check: validateATS, Run: runATS}
