package consensus

import "blockchain/block"
import "blockchain/blockchain"

func validateT_(_ *blockchain.Blockchain, b block.Block) bool {
	header := b.Header
	if header["head"] != "Test" {
		return false
	}
	return true
}

func runT_(bc *blockchain.Blockchain, b block.Block) error {
	if err := bc.AddBlock(b); err != nil {
		return err
	}
	return nil
}

var CTest = Consensus{validateT_, runT_}
