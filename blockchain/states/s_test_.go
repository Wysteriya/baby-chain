package states

import (
	"blockchain/block"
)

func validateT_(_ *StateData, b block.Block) bool {
	header := b.Header
	if header["head"] != "Test" {
		return false
	}
	return true
}

func runT_(sd *StateData, b block.Block) error {
	data := b.Data()
	(*sd)["test"] = data
	return nil
}

var STest = State{validateT_, runT_}
