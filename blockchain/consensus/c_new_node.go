package consensus

import "blockchain/block"

func validateNN(b block.Block) bool {
    return true
}

func runNN(data block.Data) error {
    return nil
}

var CNewNode Consensus = Consensus{validateNN, runNN}
