package consensus

import "blockchain/block"

type cNewNode string

func (c cNewNode) Run(block.Data) error {
    return nil
}

func (c cNewNode) Validate(b block.Block) error {
    return nil
}
