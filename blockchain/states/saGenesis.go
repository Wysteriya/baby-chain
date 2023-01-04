package states

import (
	"baby-chain/blockchain/block"
)

var SGenesis = State{
	func(_ *StateData, b block.Block) bool {
		if b.Header["head"] != "Genesis" {
			return false
		}
		return true
	},
	func(sd *StateData, b block.Block) error {
		sd.Data = b.Data
		return nil
	},
}
