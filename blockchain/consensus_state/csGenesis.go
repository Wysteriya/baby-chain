package consensus_state

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
)

var CSGenesis = ConsensusState{
	func(_ *blockchain.Blockchain, _ *StateData, b block.Block) bool {
		if b.Header[block.Head] != block.GENESIS {
			return false
		}
		return true
	},
	func(bc *blockchain.Blockchain, sd *StateData, b block.Block) error {
		if err := bc.AddBlock(b); err != nil {
			return err
		}
		// todo
		return nil
	},
}
