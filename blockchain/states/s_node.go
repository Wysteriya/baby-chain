package states

import (
	"blockchain/block"
)

func initializeNN(sd *StateData) {
	(*sd)["full_nodes"] = block.Data{}
}

func validateNN(_ *StateData, b block.Block) bool {
	header := b.Header
	if header["head"] != "NewNode" {
		return false
	}
	if _, ok := header["signature"].(string); !ok {
		return false
	}
	data := b.Data()
	if _, ok := data["public_key"].(string); !ok {
		return false
	}
	if _, ok := data["ip_address"].(string); !ok {
		return false
	}
	return true
}

func runNN(sd *StateData, b block.Block) error {
	data := b.Data()
	publicKey, _ := data["public_key"].(string)
	ipAddress, _ := data["ip_address"].(string)
	(*sd)["full_nodes"] = block.Data{publicKey: ipAddress}
	return nil
}

var SNode = State{initializeNN, validateNN, runNN}
