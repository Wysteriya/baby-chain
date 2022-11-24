package states

import (
	"blockchain/block"
)

func validateN(_ *StateData, b block.Block) bool {
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

func runN(_ *StateData, _ block.Block) error {
	return nil
}

var SNode = State{validateN, runN}
