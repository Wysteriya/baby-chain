package state

import (
	"baby-chain/blockchain/block"
)

var STest = State{
	func(_ *StateData, b block.Block) bool {
		if b.Header["head"] != "Test" {
			return false
		}
		return true
	},
	func(sd *StateData, b block.Block) error {
		sd.Data = b.Data
		return nil
	},
}
