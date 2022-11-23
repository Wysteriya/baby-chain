package consensus

import "blockchain/block"
import "blockchain/blockchain"

func validateG(bc *blockchain.Blockchain, b block.Block) bool {
    return true
}

func runG(bc *blockchain.Blockchain, b block.Block) error {
    return nil
}

var CGenesis Consensus = Consensus{validateG, runG}
