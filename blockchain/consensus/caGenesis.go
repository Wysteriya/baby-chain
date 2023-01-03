package consensus

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
)

var CGenesis = Consensus{
	func(_ *blockchain.Blockchain, b block.Block) bool {
		if b.Header["head"] != "Genesis" {
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
