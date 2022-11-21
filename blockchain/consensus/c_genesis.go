package consensus

import "blockchain/block"

type cGenesis string

func (c cGenesis) Run(block.Data) error {
    return nil
}

func (c cGenesis) Validate(b block.Block) error {
    return nil
}
