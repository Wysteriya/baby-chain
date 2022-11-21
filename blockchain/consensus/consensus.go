package consensus

import "blockchain/block"

type Consensus interface {
    Validate(block.Block) error
    Run(block.Data) error
}

type CAlgo []Consensus

func (ca *CAlgo) Check(b block.Block) error {
    for _, con := range *ca {
        if err := con.Validate(b); err != nil {
            return err
        } else if err := con.Run(b.Data()); err != nil {
            return err
        }
    }
    return nil
}

func New() CAlgo {
    return CAlgo{cGenesis}
}
