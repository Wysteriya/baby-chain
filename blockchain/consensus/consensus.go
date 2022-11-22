package consensus

import "blockchain/block"
import "blockchain/blockchain"

type Consensus struct {
    validate func(block.Block) bool
    run func(block.Data) error
}

type CAlgo []Consensus

func (ca *CAlgo) Exec(bc *blockchain.Blockchain, b block.Block) error {
    if err := bc.ValidateBlock(b, bc.Len() - 1); err != nil {
        return err
    }
    for _, con := range *ca {
        if con.validate(b) {
            if err := con.run(b.Data()); err != nil {
                return err
            }
        }
    }
    return nil
}

func New() CAlgo {
    return CAlgo{CGenesis, CNewNode}
}
