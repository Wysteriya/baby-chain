package consensus

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
)

var CTest = Consensus{
	func(_ *blockchain.Blockchain, b block.Block) bool {
		if b.Header["head"] != "Test" {
			return false
		}
		return true
	},
	func(bc *blockchain.Blockchain, b block.Block) error {
		if err := bc.AddBlock(b); err != nil {
			return err
		}
		return nil
	},
}
