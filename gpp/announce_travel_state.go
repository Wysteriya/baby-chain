package gpp

import (
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/state"
	"baby-chain/tools"
)

func initializeATS(sd *state.StateData) {
	(sd.Data)["open_announcements"] = tools.Data{}
}

func validateATS(_ *state.StateData, b block.Block) bool {
	header := b.Header
	if header["head"] != "AnnounceTravel" {
		return false
	}
	return true
}

func runATS(sd *state.StateData, b block.Block) error {
	data := b.Data
	if (sd.Data)["open_announcements"] == nil {
		initializeATS(sd)
	}
	oa, _ := (sd.Data)["open_announcements"].(tools.Data)
	hashB := tools.HashB([]byte(data.String()))
	key := (&hashB).Hex()
	oa[key] = data
	return nil
}

var SAts = state.State{Check: validateATS, Run: runATS}
