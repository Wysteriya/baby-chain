package states

import (
	"blockchain/block"
)

func initializeT_(sd *StateData) {
	(*sd)["test"] = block.Data{}
}

func validateT_(_ *StateData, b block.Block) bool {
	header := b.Header
	if header["head"] != "Test" {
		return false
	}
	return true
}

func runT_(sd *StateData, b block.Block) error {
	data := b.Data()
	test, _ := (*sd)["test"].(block.Data)
	for key, val := range data {
		test[key] = val
	}
	return nil
}

var STest = State{initializeT_, validateT_, runT_}
