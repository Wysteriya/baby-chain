package consensus

import "blockchain/block"
import "blockchain/blockchain"

func validateNN(bc *blockchain.Blockchain, b block.Block) bool {
    return true
}

func runNN(bc *blockchain.Blockchain, b block.Block) error {
    return nil
}

var CNewNode Consensus = Consensus{validateNN, runNN}
