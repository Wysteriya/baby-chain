package consensus

import "blockchain/block"

func validateG(b block.Block) bool {
    return true
}

func runG(data block.Data) error {
    return nil
}

var CGenesis Consensus = Consensus{validateG, runG}
