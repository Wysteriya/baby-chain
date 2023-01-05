package gpp

import (
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/state"
	"baby-chain/tools"
)

func validateBD(_ *state.StateData, b block.Block) bool {
	if b.Header["head"] != "Bid" {
		return false
	}
	if _, ok := b.Header["signature1"].(string); !ok {
		return false
	}
	if _, ok := b.Data["type"].(string); !ok {
		return false
	}
	if _, ok := b.Data["announcement_id"].(string); !ok {
		return false
	}
	if _, ok := b.Data["bid_amount"].(string); !ok {
		return false
	}
	if _, ok := b.Data["public_key"].(string); !ok {
		return false
	}
	return true
}

func runBD(sd *state.StateData, b block.Block) error {
	data, ok := sd.Data["OpenBids"].(tools.Data)
	if !ok {
		data = tools.Data{}
		sd.Data["OpenBids"] = data
	}
	data[b.Hash.Hex()].(tools.Data)["data"] = b.Data
	data[b.Hash.Hex()].(tools.Data)["timestamp"] = b.Timestamp.String()
	return nil
}

var SBid = state.State{Check: validateBD, Run: runBD}
