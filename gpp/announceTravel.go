package gpp

import (
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/state"
	"baby-chain/tools"
)

func validateAT(_ *state.StateData, b block.Block) bool {
	if b.Header["head"] != "AnnounceTravel" {
		return false
	}
	if _, ok := b.Header["signature1"].(string); !ok {
		return false
	}
	if _, ok := b.Data["public_key"].(string); !ok {
		return false
	}
	if _, ok := b.Data["from_lat"].(string); !ok {
		return false
	}
	if _, ok := b.Data["from_lon"].(string); !ok {
		return false
	}
	if _, ok := b.Data["to_lat"].(string); !ok {
		return false
	}
	if _, ok := b.Data["to_lon"].(string); !ok {
		return false
	}
	if _, ok := b.Data["time"].(string); !ok {
		return false
	}
	return true
}

func runAT(sd *state.StateData, b block.Block) error {
	data, ok := sd.Data["OpenAnnouncements"].(tools.Data)
	if !ok {
		data = tools.Data{}
		sd.Data["OpenAnnouncements"] = data
	}
	data[b.Hash.Hex()].(tools.Data)["data"] = b.Data
	data[b.Hash.Hex()].(tools.Data)["timestamp"] = b.Timestamp.String()
	return nil
}

var SAnnounceTravel = state.State{Check: validateAT, Run: runAT}
