package chain

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/consensus"
)

//import "blockchain/blockchain"

func validateAT(_ *blockchain.Blockchain, b block.Block) bool {
	header := b.Header
	if header["head"] != "AnnounceTravel" {
		return false
	}
	if _, ok := header["signature"].(string); !ok {
		return false
	}
	data := b.Data
	if _, ok := data["public_key"].(string); !ok {
		return false
	}
	if _, ok := data["from_lat"].(string); !ok {
		return false
	}
	if _, ok := data["from_lon"].(string); !ok {
		return false
	}
	if _, ok := data["to_lat"].(string); !ok {
		return false
	}
	if _, ok := data["to_lon"].(string); !ok {
		return false
	}
	if _, ok := data["time"].(string); !ok {
		return false
	}
	return true

}

func runAT(_ *blockchain.Blockchain, _ block.Block) error {
	return nil
}

var CAnnounceTravel = consensus.Consensus{Check: validateAT, Run: runAT}
