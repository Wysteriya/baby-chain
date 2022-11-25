package states

import (
	"blockchain/block"
)

func initializeG(_ *StateData) {
	return
}

func validateG(_ *StateData, _ block.Block) bool {
	return true
}

func runG(_ *StateData, _ block.Block) error {
	return nil
}

var SGenesis = State{initializeG, validateG, runG}
