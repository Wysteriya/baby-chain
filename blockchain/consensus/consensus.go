package consensus

import "blockchain/block"
import "blockchain/blockchain"

type Consensus struct {
    validate func(*blockchain.Blockchain, block.Block) bool
    run func(*blockchain.Blockchain, block.Block) error
}

type CAlgo []Consensus

func (ca *CAlgo) Exec(bc *blockchain.Blockchain, b block.Block) error {
    if err := bc.ValidateBlock(b, bc.Len() - 1); err != nil {
        return err
    }
    for _, con := range *ca {
        if con.validate(bc, b) {
            if err := con.run(bc, b); err != nil {
                return err
            }
        }
    }
    return nil
}

func New() CAlgo {
    return CAlgo{CGenesis, CNewNode}
}
