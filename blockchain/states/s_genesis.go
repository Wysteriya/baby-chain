package states

import (
	"blockchain/block"
)

func validateG(_ *StateData, b block.Block) bool {
	return true
}

func runG(sd *StateData, b block.Block) error {
	return nil
}

var SGenesis = State{validateG, runG}
